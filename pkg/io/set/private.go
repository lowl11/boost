package set

import "reflect"

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
