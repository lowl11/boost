package async

import (
	"context"
	"github.com/lowl11/boost/data/interfaces"
)

type group struct {
	ctx       context.Context
	semaphore interfaces.Semaphore
	tasks     []interfaces.Task
	errors    []error
}

func NewGroup(ctx context.Context) interfaces.TaskGroup {
	return &group{
		ctx:    ctx,
		tasks:  make([]interfaces.Task, 0),
		errors: make([]error, 0),
	}
}

func (g *group) Limit(limit int) interfaces.TaskGroup {
	g.semaphore = NewSemaphore(limit)
	return g
}

func (g *group) Run(f func(ctx context.Context) error) {
	t := NewTask(g.ctx)

	g.acquire()
	t.Run(func(ctx context.Context) error {
		defer g.release()

		err := f(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	g.tasks = append(g.tasks, t)
}

func (g *group) Wait() interfaces.TaskGroup {
	for _, t := range g.tasks {
		if err := t.Wait(); err != nil {
			g.errors = append(g.errors, err)
		}
	}

	g.closeSemaphore()
	return g
}

func (g *group) Errors() []error {
	return g.errors
}

func (g *group) acquire() {
	if g.semaphore == nil {
		return
	}

	g.semaphore.Acquire()
}

func (g *group) release() {
	if g.semaphore == nil {
		return
	}

	g.semaphore.Release()
}

func (g *group) closeSemaphore() {
	if g.semaphore == nil {
		return
	}

	g.semaphore.Close()
}
