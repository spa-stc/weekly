package run

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spa-stc/newsletter/internal/config"
	"github.com/spa-stc/newsletter/internal/cron"
	"github.com/spa-stc/newsletter/internal/db/postgres"
	"github.com/spa-stc/newsletter/internal/res"
	"github.com/spa-stc/newsletter/internal/services/dayupdatescron"
	"go.uber.org/zap"
)

func Run() error {
	config, err := config.NewConfig()
	if err != nil {
		return err
	}

	logger, err := zap.NewDevelopment()
	if config.Production {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return fmt.Errorf("error creating zap logger object: %s", err.Error())
	}

	db, err := postgres.NewDatabase(logger, config)
	defer db.Shutdown()
	if err != nil {
		return fmt.Errorf("error connecting to postgres database: %s", err.Error())
	}

	err = db.Migrate(res.Migrations, "migrations")
	if err != nil {
		return fmt.Errorf("error running migrations: %s", err.Error())
	}

	dayrepo := postgres.NewDayRepo(logger, db)

	cron := cron.NewRunner(logger, config)
	cron.RegisterJobs(dayupdatescron.NewDayUpdatesCronJob(logger, dayrepo, config))

	cron.Run()
	defer cron.Shutdown()

	logger.Info("startup complete, awaiting graceful shutdown message.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.Info("commencing graceful shutdown")

	return nil
}
