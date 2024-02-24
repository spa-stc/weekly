package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/spa-stc/newsletter/internal/types/models"
	"go.uber.org/zap"
)

// Postgres day repository, which holds days inside postgres.
type DayRepo struct {
	logger *zap.Logger
	db     *Database
}

func NewDayRepo(logger *zap.Logger, db *Database) *DayRepo {
	return &DayRepo{
		logger,
		db,
	}
}

func (r *DayRepo) InsertUpdateDays(ctx context.Context, days ...models.Day) error {
	err := r.db.RunInTransaction(ctx, func(tx pgx.Tx) error {
		for i := range days {
			_, err := tx.Exec(ctx, insertupdatedaysquery,
				&days[i].Year,
				&days[i].Week,
				&days[i].Day,
				&days[i].Date,
				&days[i].Lunch,
				&days[i].XPeriod,
				&days[i].RotationDay,
				&days[i].Location,
				&days[i].Notes,
				&days[i].ApInfo,
				&days[i].CCInfo,
				&days[i].Grade9,
				&days[i].Grade10,
				&days[i].Grade11,
				&days[i].Grade12,
			)

			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

const insertupdatedaysquery = `INSERT INTO newsletter.days
(year,
 week,
 day,
 actual_date,
 lunch,
 x_period,
 rotation_day,
 location, notes,
 ap_info,
 cc_info,
 grade_9,
 grade_10,
 grade_11,
 grade_12)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15)
ON CONFLICT (year, week, day) DO UPDATE
SET
actual_date = $4,
lunch = $5,
x_period = $6,
rotation_day = $7,
location = $8,
notes = $9,
ap_info = $10,
cc_info = $11,
grade_9 = $12,
grade_10 = $13,
grade_11 = $14,
grade_12 = $15,
updated_at = now();
`
