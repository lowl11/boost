package interfaces

import (
	"github.com/lowl11/boost/data/funcs"
	"time"
)

type CronRouter interface {
	Every(every int) EveryScheduler
	Cron(expression string) CronScheduler
}

type Scheduler interface {
	Action() funcs.CronHandler
	GetDuration() time.Duration
	GetStart() bool
}

type BaseScheduler interface {
	Do(jobAction funcs.CronHandler)
}

type EveryScheduler interface {
	BaseScheduler

	Milliseconds() EveryScheduler
	Seconds() EveryScheduler
	Minutes() EveryScheduler
	Hours() EveryScheduler
	Days() EveryScheduler
	Weeks() EveryScheduler
	FromStart() EveryScheduler
}

type CronScheduler interface {
	BaseScheduler
}
