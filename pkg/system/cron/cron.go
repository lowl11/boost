package cron

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/services/counter"
	"sync"
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
