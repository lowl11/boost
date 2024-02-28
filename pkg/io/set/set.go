package set

import (
	"math/rand"
	"reflect"
	"sync"
)

type Set[T any] struct {
	data         []T
	memoryLength int

	threadSafe bool
	mx         sync.Mutex
}

func New[T any](length ...int) *Set[T] {
	collection := &Set[T]{}
	if len(length) > 0 {
		collection.data = make([]T, 0, length[0])
		collection.memoryLength = length[0]
	} else {
		collection.data = make([]T, 0, 100)
		collection.memoryLength = 100
	}
	return collection
}

func (set *Set[T]) ThreadSafe() *Set[T] {
	set.threadSafe = true
	return set
}

func (set *Set[T]) Clear() *Set[T] {
	set.mx.Lock()
	defer set.mx.Unlock()
	return set._clear()
}

func (set *Set[T]) Pop() T {
	set._lock()
	defer set._unlock()
	index := set._len() - 1
	defer set._remove(index)
	return set.data[index]
}

func (set *Set[T]) Add(item T) *Set[T] {
	set._lock()
	defer set._unlock()

	if set._exist(item) {
		return set
	}

	set.data = append(set.data, item)
	return set
}

func (set *Set[T]) Remove(index int) *Set[T] {
	set._lock()
	defer set._unlock()

	if set._len() == 0 {
		return set
	}

	return set._remove(index)
}

func (set *Set[T]) Shuffle() *Set[T] {
	set._lock()
	defer set._unlock()

	for i := range set.data {
		j := rand.Intn(i + 1)
		set.data[i], set.data[j] = set.data[j], set.data[i]
	}

	return set
}

func (set *Set[T]) Slice() []T {
	set._lock()
	defer set._unlock()
	return set.data
}

func (set *Set[T]) Length() int {
	set._lock()
	defer set._unlock()
	return set._len()
}

func (set *Set[T]) _lock() {
	if !set.threadSafe {
		return
	}

	set.mx.Lock()
}

func (set *Set[T]) _unlock() {
	if !set.threadSafe {
		return
	}

	set.mx.Unlock()
}

func (set *Set[T]) _exist(searchItem T) bool {
	for _, item := range set.data {
		if reflect.DeepEqual(item, searchItem) {
			return true
		}
	}

	return false
}

func (set *Set[T]) _len() int {
	return len(set.data)
}

func (set *Set[T]) _remove(index int) *Set[T] {
	set.data = remove[T](set.data, index)
	return set
}

func (set *Set[T]) _clear() *Set[T] {
	set.data = make([]T, 0, set.memoryLength)
	return set
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
