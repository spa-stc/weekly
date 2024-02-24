package postgres

import (
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spa-stc/newsletter/internal/config"

	// Required Database Driver.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Database struct {
	logger *zap.Logger
	db     *pgxpool.Pool
	url    string
}

func NewDatabase(logger *zap.Logger, c config.Config) (*Database, error) {

	pool, err := pgxpool.New(context.Background(), c.Db.Url)
	if err != nil {
		return &Database{}, err
	}

	db := &Database{
		logger,
		pool,
		c.Db.Url,
	}

	return db, nil
}

func (db *Database) Shutdown() {
	db.logger.Info("shutting down database")
	// This Will Not Cut Off Open Connections and Transactions
	// db.db.Close()
	db.logger.Info("database shutdown complete")
}

func (db *Database) GetConnection(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := db.db.Acquire(ctx)
	if err != nil {
		db.logger.Error("error aquiring database connection", zap.Error(err))
		return conn, err
	}

	return conn, err
}

// Run A Contained Function in A Transaction, Handling Rollbacks and Commits.
func (db *Database) RunInTransaction(ctx context.Context, fn func(pgx.Tx) error) error {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		db.logger.Warn("error opening database connection", zap.Error(err))
		return err
	}
	defer func() {
		conn.Release()
	}()

	tx, err := db.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		db.logger.Error("error starting database transaction", zap.Error(err))
		return err
	}

	return db.runInTx(ctx, tx, fn)
}

// Handle Errors In The Transaction, Rolling Back if One Occurs
// TODO: Figure Out If Panic Handling Here is Necessary
// (It Was With database/sql) (We Probably Will Be Fine Without This).
func (db *Database) runInTx(ctx context.Context, tx pgx.Tx, fn func(pgx.Tx) error) error {
	defer func() {
		if err := recover(); err != nil {
			db.logger.Error("failed to rollback database transaction", zap.Any("any type", err))
			panic(err)
		}
	}()

	if err := fn(tx); err != nil {
		if rollbackError := tx.Rollback(ctx); rollbackError != nil {
			db.logger.Warn("inital tx rollback failed", zap.Error(rollbackError))
		}
		return err
	}

	return tx.Commit(ctx)
}

// Use Golang-Migrate To Apply Migrations to the Database
// Takes an embed.FS, and the path where migrations will be found in that fs.
func (db *Database) Migrate(migrations embed.FS, path string) error {
	// Use IOFS To Perform Migrations
	fs, err := iofs.New(migrations, path)
	if err != nil {
		return fmt.Errorf("error creating migrations iofs: %s", err.Error())
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", fs, db.url)
	if err != nil {
		return fmt.Errorf("error creating migratior: %s", err.Error())
	}

	if err = migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			db.logger.Info("database migrations complete with no change")
			return nil
		}

		return fmt.Errorf("error completing database migrations: %s", err.Error())
	}

	return nil
}
