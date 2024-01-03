package async

import (
	"context"
	"github.com/lowl11/boost/data/interfaces"
)

type Group struct {
	ctx       context.Context
	semaphore interfaces.Semaphore
	tasks     []interfaces.Task
	errors    []error
}

func NewGroup(ctx context.Context) *Group {
	return &Group{
		ctx:    ctx,
		tasks:  make([]interfaces.Task, 0),
		errors: make([]error, 0),
	}
}

func (group *Group) Limit(limit int) *Group {
	group.semaphore = NewSemaphore(limit)
	return group
}

func (group *Group) Run(f func(ctx context.Context) error) {
	task := NewTask(group.ctx)

	task.Run(func(ctx context.Context) error {
		group.acquire()
		defer group.release()

		return f(ctx)
	})

	group.tasks = append(group.tasks, task)
}

func (group *Group) Wait() {
	for _, task := range group.tasks {
		if err := task.Wait(); err != nil {
			group.errors = append(group.errors, err)
		}
	}

	group.closeSemaphore()
}

func (group *Group) Errors() []error {
	return group.errors
}

func (group *Group) acquire() {
	if group.semaphore == nil {
		return
	}

	group.semaphore.Acquire()
}

func (group *Group) release() {
	group.semaphore.Release()
}

func (group *Group) closeSemaphore() {
	if group.semaphore == nil {
		return
	}

	group.semaphore.Close()
}
