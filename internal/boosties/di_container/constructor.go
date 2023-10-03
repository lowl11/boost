package di_container

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/internal/containers/tqueue"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/log"
	"reflect"
)

func callValues(tq *tqueue.Queue, constructor any, services map[reflect.Type]*serviceInfo) []reflect.Value {
	constructorType := reflect.TypeOf(constructor)

	argsCount := constructorType.NumIn()
	arguments := make([]reflect.Value, 0, argsCount)

	// prepare params/arguments
	for i := 0; i < argsCount; i++ {
		var argType reflect.Value
		if constructorType.In(i).Kind() == reflect.Ptr {
			argType = reflect.New(constructorType.In(i).Elem())
		} else {
			argType = reflect.New(constructorType.In(i))
		}
		unwrappedArgType := type_helper.UnwrapValue(argType)

		// service case
		argService, isArgService := services[unwrappedArgType.Type()]
		if !isArgService {
			// check for primitives
			primitiveValue := tq.Dequeue(unwrappedArgType.Type())
			if primitiveValue == nil {
				panic("Required argument for constructor not found")
			}

			arguments = append(arguments, reflect.ValueOf(primitiveValue))
			continue
		}

		switch argService.mode {
		case di_modes.Singleton:
			// singleton arg
			if argService.instance != nil {
				return []reflect.Value{reflect.ValueOf(argService.instance)}
			}

			values := callValues(argService.tq, argService.constructor, services)
			if len(values) > 0 {
				argService.instance = values[0].Interface()
				arguments = append(arguments, values[0])
			}
		case di_modes.Scoped:
			// scoped arg (check for current request)
			values := callValues(argService.tq, argService.constructor, services)
			if len(values) == 0 {
				panic("Constructor returns anything")
			}

			arguments = append(arguments, values[0])
		case di_modes.Transient:
			// transient arg (just create new one)
			values := callValues(argService.tq, argService.constructor, services)
			if len(values) == 0 {
				panic("Constructor returns anything")
			}

			arguments = append(arguments, values[0])
		}
	}

	// return values
	return reflect.ValueOf(constructor).Call(arguments)
}

func call(tq *tqueue.Queue, constructor any, services map[reflect.Type]*serviceInfo) any {
	var result = callValues(tq, constructor, services)
	if len(result) == 0 {
		return nil
	}

	for i := 1; i < len(result)-1; i++ {
		err, isError := result[i].Interface().(error)
		if !isError {
			continue
		}

		if err != nil {
			log.Error(err, "Get instance error")
		}
	}

	// return the instance
	return result[0].Interface()
}

type constructor struct {
	// value & type
	f any
	t reflect.Type

	// checks
	kindCheck   bool
	returnCheck bool
}

func newConstructor(f any) *constructor {
	return &constructor{
		f: f,
		t: reflect.TypeOf(f),

		kindCheck:   true,
		returnCheck: true,
	}
}

func (c *constructor) IsFunc() *constructor {
	c.kindCheck = c.t.Kind() == reflect.Func
	if c.t.Kind() != reflect.Func {
		panic("Given constructor is not the func type")
	}
	return c
}

func (c *constructor) IsReturnMatch(returnType reflect.Type) *constructor {
	if c.t.NumOut() == 0 {
		c.returnCheck = false
		panic("Constructor returns anything")
	}

	realReturnType := type_helper.UnwrapType(c.t.Out(0))
	if realReturnType != returnType {
		c.returnCheck = false
		panic("Constructor return value type is not correct")
	}

	return c
}

func (c *constructor) GetDependencies() []reflect.Type {
	deps := make([]reflect.Type, 0, c.t.NumIn())
	for i := 0; i < c.t.NumIn(); i++ {
		dep := type_helper.UnwrapType(c.t.In(i))
		deps = append(deps, dep)
	}
	return deps
}
