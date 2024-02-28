package cron

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/fast_handler/counter"
	"sync"
	"time"
)

type (
	CronRouter = interfaces.CronRouter
)

type Config struct {
	ErrorHandler funcs.CronErrorHandler
}

type Cron struct {
	errorHandler      funcs.CronErrorHandler
	schedulersChannel chan interfaces.Scheduler
	schedulers        []interfaces.Scheduler
	counter           *counter.Counter

	mutex sync.Mutex
}

func New(config Config, counter *counter.Counter) *Cron {
	cron := &Cron{
		schedulersChannel: make(chan interfaces.Scheduler),
		schedulers:        make([]interfaces.Scheduler, 0),
		errorHandler:      config.ErrorHandler,
		counter:           counter,
	}

	go cron.readActions()

	return cron
}

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

func (cron *Cron) addScheduler(scheduler interfaces.Scheduler) {
	cron.mutex.Lock()
	defer cron.mutex.Unlock()

	cron.schedulers = append(cron.schedulers, scheduler)
}

func (cron *Cron) readActions() {
	for scheduler := range cron.schedulersChannel {
		cron.addScheduler(scheduler)
	}
}
