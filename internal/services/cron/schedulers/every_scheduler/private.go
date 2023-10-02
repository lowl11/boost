package every_scheduler

import "time"

func (scheduler *Scheduler) getDuration() time.Duration {
	return time.Duration(scheduler.every) * scheduler.duration
}
