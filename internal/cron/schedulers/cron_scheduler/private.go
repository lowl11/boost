package cron_scheduler

import (
	"github.com/aptible/supercronic/cronexpr"
	"time"
)

func (scheduler *Scheduler) getDuration(expression string) time.Duration {
	nextTime := cronexpr.MustParse(expression).Next(time.Now())
	difference := nextTime.Sub(time.Now())

	hours := time.Duration(difference.Hours()) * time.Hour
	minutes := time.Duration(difference.Minutes()) * time.Minute
	seconds := time.Duration(difference.Seconds()) * time.Second
	milliseconds := time.Duration(difference.Milliseconds()) * time.Millisecond

	return hours + minutes + seconds + milliseconds
}
