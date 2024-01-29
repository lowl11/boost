package cron_scheduler

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
)

type Scheduler struct {
	schedulersChannel chan interfaces.Scheduler
	cronExpression    string
	fromStart         bool

	jobAction funcs.CronHandler
}

func New(schedulersChannel chan interfaces.Scheduler, cronExpression string) *Scheduler {
	return &Scheduler{
		schedulersChannel: schedulersChannel,
		cronExpression:    cronExpression,
	}
}
