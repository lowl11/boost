package cron

import (
	"github.com/lowl11/boost/data/interfaces"
	"time"
)

func (cron *Cron) Every(every int) interfaces.EveryScheduler {
	cron.counter.CronAction()
	return newEveryScheduler(cron.schedulersChannel, every)
}

func (cron *Cron) Cron(cronExpression string) interfaces.CronScheduler {
	cron.counter.CronAction()
	return newCronScheduler(cron.schedulersChannel, cronExpression)
}

func (cron *Cron) Run() {
	close(cron.schedulersChannel)

	time.Sleep(time.Millisecond * 250)

	for _, scheduler := range cron.schedulers {
		go newRunner(scheduler).
			ErrorHandler(cron.errorHandler).
			FromStart(scheduler.GetStart()).
			StartTicker()
	}

	infinite := make(chan bool, 1)
	<-infinite
}

func (cron *Cron) RunAsync() {
	go cron.Run()
}
