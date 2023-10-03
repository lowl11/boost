package di

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/internal/boosties/di_container"
	"reflect"
)

func Get[T any](params ...any) *T {
	return di_container.Get().Get(reflect.TypeOf(new(T)), params...).(*T)
}

func AddTransient[T any](constructor any, dependencies ...any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Transient, dependencies...)
}

func AddScoped[T any](constructor any, dependencies ...any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Scoped, dependencies...)
}

func AddSingleton[T any](constructor any, dependencies ...any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Singleton, dependencies...)
}

func MapControllers(constructors ...any) {
	di_container.Get().MapControllers(constructors...)
}
