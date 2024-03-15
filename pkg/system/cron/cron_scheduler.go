package cron

import (
	"github.com/aptible/supercronic/cronexpr"
	"github.com/lowl11/boost/data/interfaces"
	"time"
)

type cronScheduler struct {
	schedulersChannel chan interfaces.Scheduler
	cronExpression    string
	fromStart         bool

	jobAction interfaces.CronHandler
}

func newCronScheduler(schedulersChannel chan interfaces.Scheduler, cronExpression string) *cronScheduler {
	return &cronScheduler{
		schedulersChannel: schedulersChannel,
		cronExpression:    cronExpression,
	}
}

func (scheduler *cronScheduler) Action() interfaces.CronHandler {
	return scheduler.jobAction
}

func (scheduler *cronScheduler) FromStart() interfaces.CronScheduler {
	scheduler.fromStart = true
	return scheduler
}

func (scheduler *cronScheduler) GetStart() bool {
	return scheduler.fromStart
}

func (scheduler *cronScheduler) GetDuration() time.Duration {
	return scheduler.getDuration(scheduler.cronExpression)
}

func (scheduler *cronScheduler) Do(jobAction interfaces.CronHandler) {
	scheduler.jobAction = jobAction
	scheduler.schedulersChannel <- scheduler
}

func (scheduler *cronScheduler) getDuration(expression string) time.Duration {
	nextTime := cronexpr.MustParse(expression).Next(time.Now())
	difference := nextTime.Sub(time.Now())

	hours := time.Duration(difference.Hours()) * time.Hour
	minutes := time.Duration(difference.Minutes()) * time.Minute
	seconds := time.Duration(difference.Seconds()) * time.Second
	milliseconds := time.Duration(difference.Milliseconds()) * time.Millisecond

	return hours + minutes + seconds + milliseconds
}
