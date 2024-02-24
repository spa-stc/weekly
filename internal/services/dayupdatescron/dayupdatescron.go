package dayupdatescron

import (
	"context"
	"time"

	"github.com/apognu/gocal"
	"github.com/spa-stc/newsletter/internal/config"
	"github.com/spa-stc/newsletter/internal/types/models"
	"github.com/spa-stc/newsletter/internal/types/repositories"
	"go.uber.org/zap"
)

// Cron Job That Handles Periodic Day Updates.
type Job struct {
	logger  *zap.Logger
	dayrepo repositories.DayRepository
	config  config.Config
}

// JobDay is a superset of day, which is used to handle running jobs.
type JobDay struct {
	// Internal Day Model
	Day models.Day

	// Time Used to Parse Info From Sheets and Ical.
	RelatedTime time.Time `json:"omit"`
}

func NewDayUpdatesCronJob(logger *zap.Logger, dayrepo repositories.DayRepository, config config.Config) *Job {
	return &Job{
		logger,
		dayrepo,
		config,
	}
}

func (d *Job) GetSpec() string {
	if d.config.Production {
		// Every day at 00:00
		return "0 0 * * *"
	}

	return "* * * * *"
}

func (d *Job) Run() {
	now := time.Now()
	// now := time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC)
	d.logger.Info("beginning day updates cron job")

	// Job Days that will be used to get the results of the job.
	jobdays := generateJobDays(now)

	// Get The Basic id Fields from the embedded date time.
	jobdays = populateJobDays(jobdays)

	// Download and use the lunch ical.
	calparser, err := getLunchIcalFromURL(d.config.IcalURL)
	if err != nil {
		d.logger.Error("error downloading or parsing lunch ical", zap.Error(err))
		d.logger.Warn("aborting cron job run")
		return
	}

	// Get Lunches From The Related ICAL
	jobdays = populateLunches(jobdays, calparser)

	csv, err := GetCSVReaderFromInfo(d.config.Google.SheetID, d.config.Google.Sheet)
	if err != nil {
		d.logger.Error("error downloading csv value", zap.Error(err))
		d.logger.Warn("aborting cron job run")
		return
	}
	defer csv.Close()

	data, err := GetCSVDataFromCSV(csv)
	if err != nil {
		d.logger.Error("error parsing csv data from ioreader", zap.Error(err))
		d.logger.Warn("aborting cron job run")
		return
	}

	// Get other information for job days.
	jobdays = populateOtherInfoForJobDays(jobdays, data)

	const contextlen = 10
	ctx, c := context.WithTimeout(context.Background(), time.Second*contextlen)
	defer c()

	// Insert Job Days At The End Of the Job.
	err = d.dayrepo.InsertUpdateDays(ctx, jobDaysToDays(jobdays)...)
	if err != nil {
		d.logger.Error("error inserting jobdays into dayrepository at end of cron job", zap.Error(err))
		return
	}

	d.logger.Info("completed cron job")
}

func jobDaysToDays(j []JobDay) []models.Day {
	var days []models.Day
	for _, v := range j {
		days = append(days, v.Day)
	}

	return days
}

// Populate the Year, Week, and Day fields.
func populateJobDays(j []JobDay) []JobDay {
	var newjs []JobDay
	for _, v := range j {
		year, week, day := getYearWeekDay(v.RelatedTime)
		v.Day.Day = day
		v.Day.Year = year
		v.Day.Week = week
		v.Day.Date = v.RelatedTime

		newjs = append(newjs, v)
	}

	return newjs
}

// Populate the Lunch Field of a JobDay.
func populateLunches(j []JobDay, parser *gocal.Gocal) []JobDay {
	var newjs []JobDay
	for _, v := range j {
		v.Day.Lunch = GetLunch(parser, v.RelatedTime)
		newjs = append(newjs, v)
	}

	return newjs
}
