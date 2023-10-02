package every_scheduler

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"time"
)

func (scheduler *Scheduler) Milliseconds() interfaces.EveryScheduler {
	scheduler.duration = time.Millisecond
	return scheduler
}

func (scheduler *Scheduler) Seconds() interfaces.EveryScheduler {
	scheduler.duration = time.Second
	return scheduler
}

func (scheduler *Scheduler) Minutes() interfaces.EveryScheduler {
	scheduler.duration = time.Minute
	return scheduler
}

func (scheduler *Scheduler) Hours() interfaces.EveryScheduler {
	scheduler.duration = time.Hour
	return scheduler
}

func (scheduler *Scheduler) Days() interfaces.EveryScheduler {
	scheduler.duration = time.Hour * 24
	return scheduler
}

func (scheduler *Scheduler) Weeks() interfaces.EveryScheduler {
	scheduler.duration = time.Hour * 24 * 7
	return scheduler
}

func (scheduler *Scheduler) FromStart() interfaces.EveryScheduler {
	scheduler.fromStart = true
	return scheduler
}

func (scheduler *Scheduler) GetStart() bool {
	return scheduler.fromStart
}

func (scheduler *Scheduler) Do(jobAction funcs.CronHandler) {
	scheduler.jobAction = jobAction
	scheduler.schedulersChannel <- scheduler
}

func (scheduler *Scheduler) Action() funcs.CronHandler {
	return scheduler.jobAction
}

func (scheduler *Scheduler) GetDuration() time.Duration {
	return scheduler.getDuration()
}
