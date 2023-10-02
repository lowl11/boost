package di

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/internal/boosties/di_container"
	"reflect"
)

func Get[T any](params ...any) *T {
	return di_container.Get().Get(reflect.TypeOf(new(T)), params...).(*T)
}

func AddTransient[T any](constructor any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Transient)
}

func AddScoped[T any](constructor any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Scoped)
}

func AddSingleton[T any](constructor any) {
	di_container.Get().Register(reflect.TypeOf(new(T)), constructor, di_modes.Singleton)
}
