package di

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/internal/di_container"
	"reflect"
)

func Get[T any]() *T {
	object := di_container.Get().Get(reflect.TypeOf(new(T)))
	if object == nil {
		return nil
	}

	typed, ok := object.(*T)
	if !ok {
		return nil
	}

	return typed
}

func Interface[T any]() T {
	object := di_container.Get().Get(reflect.TypeOf(new(T)))
	if object == nil {
		return *new(T)
	}

	typed, ok := object.(T)
	if !ok {
		return *new(T)
	}

	return typed
}

func RegisterTransient[T any](constructor any, dependencies ...any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Transient, dependencies...)
}

// Register registers object in "singleton" mode
func Register[T any](constructor any, dependencies ...any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Singleton, dependencies...)
}

func MapControllers(constructors ...any) {
	di_container.Get().MapControllers(constructors...)
}
