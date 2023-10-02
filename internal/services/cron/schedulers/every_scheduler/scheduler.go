package every_scheduler

import (
	"github.com/lowl11/boost/data/funcs"
	"github.com/lowl11/boost/data/interfaces"
	"time"
)

type Scheduler struct {
	schedulersChannel chan interfaces.Scheduler
	every             int
	fromStart         bool

	duration time.Duration

	jobAction funcs.CronHandler
}

func New(schedulersChannel chan interfaces.Scheduler, every int) *Scheduler {
	return &Scheduler{
		schedulersChannel: schedulersChannel,
		every:             every,
	}
}
