package queue

import "reflect"

type Queue struct {
	dataType reflect.Type
	data     []any
	len      int
}

func New(t reflect.Type) *Queue {
	return &Queue{
		dataType: t,
		data:     make([]any, 0, 10),
	}
}
