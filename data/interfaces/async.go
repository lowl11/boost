package interfaces

import "context"

type Semaphore interface {
	Acquire()
	Release()
}

type Task interface {
	Run(f func(ctx context.Context) error) Task
	Wait() error
	IsDone() bool
	Error() error
}
