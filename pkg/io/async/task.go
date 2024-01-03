package async

import (
	"context"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/io/exception"
	"sync"
	"sync/atomic"
)

type task struct {
	wg *sync.WaitGroup

	ctx    context.Context
	cancel func()

	errOnce sync.Once
	err     error

	isDone atomic.Bool
}

func NewTask(ctx context.Context) interfaces.Task {
	ctx, cancel := context.WithCancel(ctx)
	return &task{
		ctx:    ctx,
		cancel: cancel,
		wg:     &sync.WaitGroup{},
	}
}

func (t *task) Run(f func(ctx context.Context) error) interfaces.Task {
	t.wg.Add(1)

	go func() {
		defer t.isDone.Store(true)
		defer t.wg.Done()

		if err := exception.Try(func() error {
			return f(t.ctx)
		}); err != nil {
			t.errOnce.Do(func() {
				t.err = err

				if t.cancel != nil {
					t.cancel()
				}
			})
		}
	}()

	return t
}

func (t *task) Wait() error {
	t.wg.Wait()

	if t.cancel != nil {
		t.cancel()
	}

	return t.err
}

func (t *task) IsDone() bool {
	return t.isDone.Load()
}

func (t *task) Error() error {
	return t.err
}
