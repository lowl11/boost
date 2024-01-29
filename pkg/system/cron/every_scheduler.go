package cron

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"time"
)

type everyScheduler struct {
	schedulersChannel chan interfaces.Scheduler
	every             int
	fromStart         bool

	duration time.Duration

	jobAction funcs.CronHandler
}

func newEveryScheduler(schedulersChannel chan interfaces.Scheduler, every int) *everyScheduler {
	return &everyScheduler{
		schedulersChannel: schedulersChannel,
		every:             every,
	}
}

func (scheduler *everyScheduler) Milliseconds() interfaces.EveryScheduler {
	scheduler.duration = time.Millisecond
	return scheduler
}

func (scheduler *everyScheduler) Seconds() interfaces.EveryScheduler {
	scheduler.duration = time.Second
	return scheduler
}

func (scheduler *everyScheduler) Minutes() interfaces.EveryScheduler {
	scheduler.duration = time.Minute
	return scheduler
}

func (scheduler *everyScheduler) Hours() interfaces.EveryScheduler {
	scheduler.duration = time.Hour
	return scheduler
}

func (scheduler *everyScheduler) Days() interfaces.EveryScheduler {
	scheduler.duration = time.Hour * 24
	return scheduler
}

func (scheduler *everyScheduler) Weeks() interfaces.EveryScheduler {
	scheduler.duration = time.Hour * 24 * 7
	return scheduler
}

func (scheduler *everyScheduler) FromStart() interfaces.EveryScheduler {
	scheduler.fromStart = true
	return scheduler
}

func (scheduler *everyScheduler) GetStart() bool {
	return scheduler.fromStart
}

func (scheduler *everyScheduler) Do(jobAction funcs.CronHandler) {
	scheduler.jobAction = jobAction
	scheduler.schedulersChannel <- scheduler
}

func (scheduler *everyScheduler) Action() funcs.CronHandler {
	return scheduler.jobAction
}

func (scheduler *everyScheduler) GetDuration() time.Duration {
	return scheduler.getDuration()
}

func (scheduler *everyScheduler) getDuration() time.Duration {
	return time.Duration(scheduler.every) * scheduler.duration
}
