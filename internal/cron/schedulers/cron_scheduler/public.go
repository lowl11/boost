package cron_scheduler

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"time"
)

func (scheduler *Scheduler) Action() funcs.CronHandler {
	return scheduler.jobAction
}

func (scheduler *Scheduler) FromStart() interfaces.CronScheduler {
	scheduler.fromStart = true
	return scheduler
}

func (scheduler *Scheduler) GetStart() bool {
	return scheduler.fromStart
}

func (scheduler *Scheduler) GetDuration() time.Duration {
	return scheduler.getDuration(scheduler.cronExpression)
}

func (scheduler *Scheduler) Do(jobAction funcs.CronHandler) {
	scheduler.jobAction = jobAction
	scheduler.schedulersChannel <- scheduler
}
