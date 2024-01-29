package cron

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
	"time"
)

type runner struct {
	scheduler    interfaces.Scheduler
	errorHandler funcs.CronErrorHandler
	fromStart    bool
}

func newRunner(scheduler interfaces.Scheduler) *runner {
	return &runner{
		scheduler: scheduler,
	}
}

func (runner *runner) ErrorHandler(handler funcs.CronErrorHandler) *runner {
	runner.errorHandler = handler
	return runner
}

func (runner *runner) StartTicker() {
	if runner.fromStart {
		go func() {
			time.Sleep(time.Millisecond * 500)
			runner.runAction()
		}()
	}

	for {
		ticker := time.NewTicker(runner.scheduler.GetDuration())

		<-ticker.C

		runner.runAction()
		ticker.Reset(runner.scheduler.GetDuration())
	}
}

func (runner *runner) FromStart(fromStart bool) *runner {
	runner.fromStart = fromStart
	return runner
}

func (runner *runner) runAction() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(exception.CatchPanic(err), "PANIC RECOVERED")
		}
	}()

	jobAction := runner.scheduler.Action()

	if err := jobAction(); err != nil {
		if runner.errorHandler != nil {
			if innerError := runner.errorHandler(err); innerError != nil {
				log.Error(innerError, "Cron error handler error")
			}
		}

		log.Error(err, "Execute job action error")
	}
}