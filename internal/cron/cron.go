package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spa-stc/newsletter/internal/config"
	"go.uber.org/zap"
)

// Cron service definition.
type Runner struct {
	cron   *cron.Cron
	logger *zap.Logger
}

// A Registerable Cron Job Service.
type Job interface {
	// Get the Cron Spec Of the Job.
	GetSpec() string

	// Actually Run the Job.
	Run()
}

// Get a new cron runner.
func NewRunner(logger *zap.Logger, _ config.Config) *Runner {
	runner := cron.New(cron.WithLocation(time.FixedZone("CST", -5)))

	return &Runner{
		runner,
		logger,
	}
}

// Run the cron service in a separite goroutine.
func (r *Runner) Run() {
	r.logger.Info("running cron service")
	go r.cron.Run()
}

// Stop the cron service after 10 seconds.
func (r *Runner) Shutdown() {
	r.logger.Info("shutting down cron service")
	ctx := r.cron.Stop()
	<-ctx.Done()
	r.logger.Info("successfully shut down cron service")
}

// Register Cron Jobs To Be Run By The Scheduler.
func (r *Runner) RegisterJobs(jobs ...Job) error {
	for _, v := range jobs {
		_, err := r.cron.AddFunc(v.GetSpec(), v.Run)
		if err != nil {
			return fmt.Errorf("error registering cron jobs: %s", err.Error())
		}
	}

	return nil
}
