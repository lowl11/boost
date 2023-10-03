package queue

import (
	"github.com/lowl11/boost/internal/helpers/container_helper"
	"reflect"
)

func (queue *Queue) Enqueue(value any) *Queue {
	queue.data = append(queue.data, value)
	queue.copyData = append(queue.copyData, value)
	queue.len++
	return queue
}

func (queue *Queue) Dequeue() any {
	if len(queue.data) == 0 {
		if len(queue.copyData) == 0 {
			return nil
		}

		copy(queue.data, queue.copyData)
	}

	first := queue.data[0]
	queue.data = container_helper.Remove(queue.data, 0)
	queue.len--
	if queue.len == 0 {
		copy(queue.data, queue.copyData)
	}
	return first
}

func (queue *Queue) Len() int {
	return queue.len
}

func (queue *Queue) Copy() *Queue {
	cp := New(queue.dataType)
	if len(queue.data) == 0 && len(queue.copyData) > 0 {
		queue.data = queue.copyData
	}
	for _, item := range queue.data {
		cp.data = append(cp.data, item)
	}
	return cp
}

func (queue *Queue) DT() reflect.Type {
	return queue.dataType
}
