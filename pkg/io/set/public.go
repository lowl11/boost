package set

import "math/rand"

func (set *Set[T]) ThreadSafe() *Set[T] {
	set.threadSafe = true
	return set
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
