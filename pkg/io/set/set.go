package set

import "sync"

type Set[T any] struct {
	data []T

	threadSafe bool
	mx         sync.Mutex
}

func New[T any](length ...int) *Set[T] {
	collection := &Set[T]{}
	if len(length) > 0 {
		collection.data = make([]T, 0, length[0])
	} else {
		collection.data = make([]T, 0, 100)
	}
	return collection
}
