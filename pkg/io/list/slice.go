package list

import (
	"github.com/lowl11/boost/pkg/io/types"
	"math/rand"
	"sort"
)

func All[T any](source []T, fn func(T) bool) bool {
	for _, element := range source {
		if !fn(element) {
			return false
		}
	}

	return true
}

func Any[T any](source []T, fn func(T) bool) bool {
	for _, element := range source {
		if fn(element) {
			return true
		}
	}

	return false
}

func Each[T any](source []T, fn func(T)) {
	for _, element := range source {
		fn(element)
	}
}

func Filter[T any](source []T, fn func(T) bool) []T {
	dst := make([]T, 0, len(source))
	for _, element := range source {
		if fn(element) {
			dst = append(dst, element)
		}
	}
	return dst
}

func FilterNot[T any](source []T, fn func(T) bool) []T {
	dst := make([]T, 0, len(source))
	for _, element := range source {
		if !fn(element) {
			dst = append(dst, element)
		}
	}
	return dst
}

func Single[T any](source []T, fn func(T) bool) *T {
	for _, element := range source {
		if fn(element) {
			return &element
		}
	}

	return nil
}

func Map[T any, U any](source []T, fn func(T) U) []U {
	newSlice := make([]U, len(source))
	for i, element := range source {
		newSlice[i] = fn(element)
	}
	return newSlice
}

func Reverse[T any](source []T) []T {
	length := len(source)
	out := make([]T, length)
	copy(out, source)
	for i := 0; i < length/2; i++ {
		out[i], out[length-i-1] = out[length-i-1], out[i]
	}
	return out
}

func Shuffle[T any](source []T, r ...rand.Source) []T {
	out := make([]T, len(source))
	copy(out, source)
	if r == nil || len(r) == 0 {
		rand.Shuffle(len(out), func(i, j int) {
			out[i], out[j] = out[j], out[i]
		})

		return out
	}

	rnd := rand.New(r[0])
	rnd.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out
}

func Sort[T any](source []T, less func(a, b T) bool) []T {
	if len(source) <= 1 {
		return source
	}

	out := make([]T, len(source))
	copy(out, source)
	sort.Slice(out, func(i, j int) bool {
		return less(out[i], out[j])
	})
	return out
}

func Add[T any](source []T, elements ...T) []T {
	return Set(source, len(source), elements...)
}

func AddLeft[T any](source []T, elements ...T) []T {
	return Set(source, 0, elements...)
}

func Set[T any](source []T, index int, elements ...T) []T {
	if index < 0 {
		index = 0
	}

	if index >= len(source) {
		return append(source, elements...)
	}

	return append(source[:index], append(elements, source[index:]...)...)
}

func Remove[T any](source []T, index ...int) []T {
	if len(index) == 0 {
		return source
	}

	if len(index) == 1 {
		i := index[0]

		if i < 0 || i >= len(source) {
			return source
		}

		return append(source[:i], source[i+1:]...)
	}

	sort.Ints(index)

	dst := make([]T, 0, len(source))

	prev := 0
	for _, i := range index {
		if i < 0 || i >= len(source) {
			continue
		}

		dst = append(dst, source[prev:i]...)
		prev = i + 1
	}

	return append(dst, source[prev:]...)
}

func SliceAny[T any](source []T, fn ...func(T) any) []any {
	var mapFn func(T) any
	if len(fn) > 0 {
		mapFn = fn[0]
	}

	sliceAny := make([]any, len(source))
	for i, element := range source {
		if mapFn != nil {
			sliceAny[i] = mapFn(element)
		} else {
			sliceAny[i] = element
		}
	}
	return sliceAny
}

func SliceString[T any](source []T, fn ...func(T) string) []string {
	var mapFn func(T) string
	if len(fn) > 0 {
		mapFn = fn[0]
	}

	sliceString := make([]string, len(source))
	for i, element := range source {
		if mapFn != nil {
			sliceString[i] = mapFn(element)
		} else {
			sliceString[i] = types.String(element)
		}
	}
	return sliceString
}

func Sub[T any](source []T, start, end int) []T {
	sub := make([]T, 0)
	if start < 0 || end < 0 {
		return sub
	}

	if start >= end {
		return sub
	}

	length := len(source)
	if start < length {
		if end <= length {
			sub = source[start:end]
		} else {
			zeroArray := make([]T, end-length)
			sub = append(source[start:length], zeroArray[:]...)
		}
	} else {
		zeroArray := make([]T, end-start)
		sub = zeroArray[:]
	}

	return sub
}
