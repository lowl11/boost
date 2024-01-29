package runner

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
)

type Runner struct {
	scheduler    interfaces.Scheduler
	errorHandler funcs.CronErrorHandler
	fromStart    bool
}

func New(scheduler interfaces.Scheduler) *Runner {
	return &Runner{
		scheduler: scheduler,
	}
}
