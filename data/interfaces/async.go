package interfaces

import "context"

type Semaphore interface {
	Acquire()
	Release()
	Close()
}

type Task interface {
	Run(f func(ctx context.Context) error) Task
	Wait() error
	IsDone() bool
	Error() error
}

type TaskGroup interface {
	Limit(limit int) TaskGroup
	Run(f func(ctx context.Context) error)
	Wait()
	Errors() []error
}
