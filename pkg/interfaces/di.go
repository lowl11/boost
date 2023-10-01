package interfaces

import "reflect"

type DependencyContainer interface {
	Register(t reflect.Type, constructor any, mode int)
	Get(t reflect.Type, params ...any) any
}
