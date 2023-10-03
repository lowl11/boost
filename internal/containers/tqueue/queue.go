package tqueue

import (
	"github.com/lowl11/boost/internal/containers/queue"
	"reflect"
)

type Queue struct {
	queues map[reflect.Type]*queue.Queue
}

func New() *Queue {
	return &Queue{
		queues: make(map[reflect.Type]*queue.Queue),
	}
}
