package async

import (
	"context"
	"github.com/lowl11/boost/pkg/io/exception"
	"sync"
	"sync/atomic"
)

type Task struct {
	wg *sync.WaitGroup

	ctx    context.Context
	cancel func()

	errOnce sync.Once
	err     error

	isDone atomic.Bool
}

func NewTask(ctx context.Context) *Task {
	ctx, cancel := context.WithCancel(ctx)
	return &Task{
		ctx:    ctx,
		cancel: cancel,
		wg:     &sync.WaitGroup{},
	}
}

func (task *Task) Run(f func(ctx context.Context) error) {
	task.wg.Add(1)

	go func() {
		defer task.isDone.Store(true)
		defer task.wg.Done()

		if err := exception.Try(func() error {
			return f(task.ctx)
		}); err != nil {
			task.errOnce.Do(func() {
				task.err = err

				if task.cancel != nil {
					task.cancel()
				}
			})
		}
	}()
}

func (task *Task) Wait() error {
	task.wg.Wait()

	if task.cancel != nil {
		task.cancel()
	}

	return task.err
}

func (task *Task) IsDone() bool {
	return task.isDone.Load()
}

func (task *Task) Error() error {
	return task.err
}
