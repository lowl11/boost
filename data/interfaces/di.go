package interfaces

import "reflect"

type DependencyContainer interface {
	SetControllerInterface(controllerInterfaceType reflect.Type)
	SetAppType(appType reflect.Type)
	Register(t reflect.Type, constructor any, mode int, dependencies ...any)
	RegisterImplementation(impl any)
	MapControllers(constructors ...any)
	Get(t reflect.Type) any
	Check() error
}
