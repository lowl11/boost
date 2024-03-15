package interfaces

import (
	"time"
)

type CronHandler func() error
type CronErrorHandler func(err error) error

type CronRouter interface {
	Every(every int) EveryScheduler
	Cron(expression string) CronScheduler
}

type Scheduler interface {
	Action() CronHandler
	GetDuration() time.Duration
	GetStart() bool
}

type BaseScheduler interface {
	Do(jobAction CronHandler)
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
