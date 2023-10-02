package funcs

type CronHandler func() error
type CronErrorHandler func(err error) error
