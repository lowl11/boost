package interfaces

import "reflect"

type DependencyContainer interface {
	Register(t reflect.Type, constructor any, mode int, dependencies ...any)
	Get(t reflect.Type, params ...any) any
	Check() error
}
