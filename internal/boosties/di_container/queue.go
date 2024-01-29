package di_container

import (
	"reflect"
)

type queue struct {
	dataType reflect.Type
	data     []any
	copyData []any
	len      int
}

func newQueue(t reflect.Type) *queue {
	return &queue{
		dataType: t,
		data:     make([]any, 0, 10),
		copyData: make([]any, 0, 10),
	}
}

func (q *queue) Enqueue(value any) *queue {
	q.data = append(q.data, value)
	q.copyData = append(q.copyData, value)
	q.len++
	return q
}

func (q *queue) Dequeue() any {
	if len(q.data) == 0 {
		if len(q.copyData) == 0 {
			return nil
		}

		copy(q.data, q.copyData)
	}

	first := q.data[0]
	q.data = remove(q.data, 0)
	q.len--
	if q.len == 0 {
		copy(q.data, q.copyData)
	}
	return first
}

func (q *queue) Len() int {
	return q.len
}

func (q *queue) Copy() *queue {
	cp := newQueue(q.dataType)
	if len(q.data) == 0 && len(q.copyData) > 0 {
		q.data = q.copyData
	}
	for _, item := range q.data {
		cp.data = append(cp.data, item)
	}
	return cp
}

func (q *queue) DT() reflect.Type {
	return q.dataType
}

func remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
