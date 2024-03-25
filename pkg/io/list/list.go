package list

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"reflect"
	"strconv"
	"unsafe"
)

type OfSlice[T any] interface {
	All(fn func(T) bool) bool
	Any(fn func(T) bool) bool
	Each(fn func(int, T)) OfSlice[T]

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
	Exist(fn func(T) bool) bool
	Get(index int) *T
	Slice() []T
	SliceAny(fn ...func(T) any) []any
	Sub(start, end int) OfSlice[T]
	Map(fn func(T) T) []T

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
	return newOf(Filter(os.source, fn))
}

func (os *ofSlice[T]) FilterNot(fn func(T) bool) OfSlice[T] {
	return newOf(FilterNot(os.source, fn))
}

func (os *ofSlice[T]) Each(fn func(int, T)) OfSlice[T] {
	Each(os.source, fn)
	return os
}

func (os *ofSlice[T]) Single(fn func(T) bool) *T {
	return Single(os.source, fn)
}

func (os *ofSlice[T]) Exist(fn func(T) bool) bool {
	return Single(os.source, fn) != nil
}

func (os *ofSlice[T]) Get(index int) *T {
	return Get(os.source, index)
}

func (os *ofSlice[T]) Reverse() OfSlice[T] {
	return newOf(Reverse(os.source))
}

func (os *ofSlice[T]) Shuffle(source ...rand.Source) OfSlice[T] {
	return newOf(Shuffle(os.source, source...))
}

func (os *ofSlice[T]) Sort(less func(a, b T) bool) OfSlice[T] {
	return newOf(Sort(os.source, less))
}

func (os *ofSlice[T]) Add(elements ...T) OfSlice[T] {
	return newOf(Add(os.source, elements...))
}

func (os *ofSlice[T]) AddLeft(elements ...T) OfSlice[T] {
	return newOf(AddLeft(os.source, elements...))
}

func (os *ofSlice[T]) Set(index int, elements ...T) OfSlice[T] {
	return newOf(Set(os.source, index, elements...))
}

func (os *ofSlice[T]) Remove(index ...int) OfSlice[T] {
	return newOf(Remove(os.source, index...))
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

func (os *ofSlice[T]) Sub(start, end int) OfSlice[T] {
	return newOf(Sub(os.source, start, end))
}

func (os *ofSlice[T]) Map(fn func(T) T) []T {
	return Map(os.source, fn)
}

func (os *ofSlice[T]) String() string {
	return toString(os.source, false)
}

func toString(anyValue any, memory bool) string {
	if anyValue == nil {
		return ""
	}

	// string type
	if stringValue, isStr := anyValue.(string); isStr {
		return stringValue
	}

	// pointer string type
	if ptrStringValue, isPtr := anyValue.(*string); isPtr {
		return *ptrStringValue
	}

	// try cast to error
	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	// try cast to bytes
	if bytesBuffer, ok := anyValue.([]byte); ok {
		return *(*string)(unsafe.Pointer(&bytesBuffer))
	}

	// try get Stringer interface
	if stringer, ok := anyValue.(fmt.Stringer); ok {
		return stringer.String()
	}

	// try cast uuid
	if uuidValue, isUUID := anyValue.(uuid.UUID); isUUID {
		return uuidValue.String()
	} else if uuidPtr, isUidPtr := anyValue.(*uuid.UUID); isUidPtr {
		return uuidPtr.String()
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.String:
		return anyValue.(string)
	case reflect.Bool:
		return strconv.FormatBool(anyValue.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		valueInBytes, err := json.Marshal(anyValue)
		if err != nil {
			return ""
		}
		return string(valueInBytes)
	case reflect.Ptr:
		if memory || value.IsZero() {
			return fmt.Sprintf("%v", value)
		}

		return toString(value.Elem().Interface(), true)
	default:
		return fmt.Sprintf("%v", value)
	}
}
