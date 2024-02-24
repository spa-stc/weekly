package models

import "time"

// This model represents a day in the database.
type Day struct {
	Year, Week, Day int
	Date            time.Time

	Lunch,
	XPeriod,
	RotationDay,
	Location,
	Notes,
	ApInfo,
	CCInfo,
	Grade9,
	Grade10,
	Grade11,
	Grade12 string

	CreatedAt time.Time
	UpdatedAt time.Time
}
