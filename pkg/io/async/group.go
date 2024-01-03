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
	task := NewTask(g.ctx)

	task.Run(func(ctx context.Context) error {
		g.acquire()
		defer g.release()

		return f(ctx)
	})

	g.tasks = append(g.tasks, task)
}

func (g *group) Wait() interfaces.TaskGroup {
	for _, task := range g.tasks {
		if err := task.Wait(); err != nil {
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
