package cron

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"sync"
)

type Config struct {
	ErrorHandler funcs.CronErrorHandler
}

type Cron struct {
	errorHandler      funcs.CronErrorHandler
	schedulersChannel chan interfaces.Scheduler
	schedulers        []interfaces.Scheduler

	mutex sync.Mutex
}

func New(config Config) *Cron {
	cron := &Cron{
		schedulersChannel: make(chan interfaces.Scheduler),
		schedulers:        make([]interfaces.Scheduler, 0),
		errorHandler:      config.ErrorHandler,
	}

	go cron.readActions()

	return cron
}

type (
	CronRouter = interfaces.CronRouter
)

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
