package repositories

import (
	"context"

	"github.com/spa-stc/newsletter/internal/types/models"
)

// Interface defining how days are placed in the database.
type DayRepository interface {
	/*
		Insert or Update any amount of days.
		since we don't care about the difference
		in the cron service.

		This should however, respect the continuity of updated_at by re-setting it upon update
	*/
	InsertUpdateDays(ctx context.Context, days ...models.Day) error
}
