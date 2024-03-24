package list

import (
	"fmt"
	"github.com/lowl11/boost/pkg/io/types"
	"math/rand"
)

type OfSlice[T any] interface {
	All(fn func(T) bool) bool
	Any(fn func(T) bool) bool
	Each(fn func(T)) OfSlice[T]

	Filter(fn func(T) bool) OfSlice[T]
	FilterNot(fn func(T) bool) OfSlice[T]
	Reverse() OfSlice[T]
	Shuffle(source ...rand.Source) OfSlice[T]
	Sort(less func(a, b T) bool) OfSlice[T]

	Add(elements ...T) OfSlice[T]
	AddLeft(elements ...T) OfSlice[T]
	Set(index int, elements ...T) OfSlice[T]
	Remove(index ...int) OfSlice[T]
	Clear(capacity ...int) OfSlice[T]

	Single(fn func(T) bool) *T
	Slice() []T
	SliceAny(fn ...func(T) any) []any
	SliceString(fn ...func(T) string) []string
	Sub(start, end int) OfSlice[T]

	fmt.Stringer
}

func Of[T any](list []T) OfSlice[T] {
	return newOf[T](list)
}

type ofSlice[T any] struct {
	source []T
}

func newOf[T any](source []T) *ofSlice[T] {
	return &ofSlice[T]{
		source: source,
	}
}

func (os *ofSlice[T]) All(fn func(T) bool) bool {
	return All(os.source, fn)
}

func (os *ofSlice[T]) Any(fn func(T) bool) bool {
	return Any(os.source, fn)
}

func (os *ofSlice[T]) Filter(fn func(T) bool) OfSlice[T] {
	os.source = Filter(os.source, fn)
	return os
}

func (os *ofSlice[T]) FilterNot(fn func(T) bool) OfSlice[T] {
	os.source = FilterNot(os.source, fn)
	return os
}

func (os *ofSlice[T]) Each(fn func(T)) OfSlice[T] {
	Each(os.source, fn)
	return os
}

func (os *ofSlice[T]) Single(fn func(T) bool) *T {
	return Single(os.source, fn)
}

func (os *ofSlice[T]) Reverse() OfSlice[T] {
	Reverse(os.source)
	return os
}

func (os *ofSlice[T]) Shuffle(source ...rand.Source) OfSlice[T] {
	Shuffle(os.source, source...)
	return os
}

func (os *ofSlice[T]) Sort(less func(a, b T) bool) OfSlice[T] {
	Sort(os.source, less)
	return os
}

func (os *ofSlice[T]) Add(elements ...T) OfSlice[T] {
	os.source = Add(os.source, elements...)
	return os
}

func (os *ofSlice[T]) AddLeft(elements ...T) OfSlice[T] {
	os.source = AddLeft(os.source, elements...)
	return os
}

func (os *ofSlice[T]) Set(index int, elements ...T) OfSlice[T] {
	os.source = Set(os.source, index, elements...)
	return os
}

func (os *ofSlice[T]) Remove(index ...int) OfSlice[T] {
	os.source = Remove(os.source, index...)
	return os
}

func (os *ofSlice[T]) Clear(capacity ...int) OfSlice[T] {
	newCapacity := 0
	if len(capacity) > 0 {
		newCapacity = capacity[0]
	}

	os.source = make([]T, 0, newCapacity)
	return os
}

func (os *ofSlice[T]) Slice() []T {
	return os.source
}

func (os *ofSlice[T]) SliceAny(fn ...func(T) any) []any {
	return SliceAny(os.source, fn...)
}

func (os *ofSlice[T]) SliceString(fn ...func(T) string) []string {
	return SliceString(os.source, fn...)
}

func (os *ofSlice[T]) Sub(start, end int) OfSlice[T] {
	os.source = Sub(os.source, start, end)
	return os
}

func (os *ofSlice[T]) String() string {
	return types.String(os.source)
}
