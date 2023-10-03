package tqueue

import (
	"github.com/lowl11/boost/internal/containers/queue"
	"reflect"
)

func (tq *Queue) Enqueue(value any) *Queue {
	valueType := reflect.TypeOf(value)
	_, ok := tq.queues[valueType]
	if !ok {
		tq.queues[valueType] = queue.New(valueType)
	}

	tq.queues[valueType].Enqueue(value)
	return tq
}

func (tq *Queue) Dequeue(t reflect.Type) any {
	typeQueue, exist := tq.queues[t]
	if !exist {
		return nil
	}

	return typeQueue.Dequeue()
}

func (tq *Queue) Len() int {
	var length int
	for _, q := range tq.queues {
		length += q.Len()
	}
	return length
}

func (tq *Queue) Copy() *Queue {
	cp := New()
	for _, q := range tq.queues {
		cp.queues[q.DT()] = q.Copy()
	}
	return cp
}
