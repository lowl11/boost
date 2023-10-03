package queue

import (
	"github.com/lowl11/boost/internal/helpers/container_helper"
	"reflect"
)

func (queue *Queue) Enqueue(value any) *Queue {
	queue.data = append(queue.data, value)
	queue.len++
	return queue
}

func (queue *Queue) Dequeue() any {
	if len(queue.data) == 0 {
		return nil
	}

	first := queue.data[0]
	queue.data = container_helper.Remove(queue.data, 0)
	queue.len--
	return first
}

func (queue *Queue) Len() int {
	return queue.len
}

func (queue *Queue) Copy() *Queue {
	cp := New(queue.dataType)
	for _, item := range queue.data {
		cp.data = append(cp.data, item)
	}
	return cp
}

func (queue *Queue) DT() reflect.Type {
	return queue.dataType
}
