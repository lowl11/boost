package runner

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/internal/boosties/panicer"
	"github.com/lowl11/boost/log"
	"time"
)

func (runner *Runner) ErrorHandler(handler funcs.CronErrorHandler) *Runner {
	runner.errorHandler = handler
	return runner
}

func (runner *Runner) StartTicker() {
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

func (runner *Runner) FromStart(fromStart bool) *Runner {
	runner.fromStart = fromStart
	return runner
}

func (runner *Runner) runAction() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(panicer.Handle(err), "PANIC RECOVERED")
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
