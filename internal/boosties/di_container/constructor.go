package di_container

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/log"
	"reflect"
)

func call(constructor any, services map[reflect.Type]*serviceInfo) any {
	constructorType := reflect.TypeOf(constructor)

	argsCount := constructorType.NumIn()
	arguments := make([]reflect.Value, 0, argsCount)

	// prepare params/arguments
	for i := 0; i < argsCount; i++ {
		argType := reflect.New(constructorType.In(i).Elem())

		// primitive case
		//

		// service case
		argService, isArgService := services[argType.Type()]
		if !isArgService {
			//
		}

		switch argService.mode {
		case di_modes.Singleton:
			// singleton arg
		case di_modes.Scoped:
			// scoped arg (check for current request)
		case di_modes.Transient:
			// transient arg (just create new one)
		}

		arguments = append(arguments, reflect.New(constructorType.In(i).Elem()))
	}

	var result = reflect.ValueOf(constructor).Call(arguments)
	if len(result) == 0 {
		return nil
	}

	if len(result) > 1 {
		for i := 1; i < len(result)-1; i++ {
			err, isError := result[i].Interface().(error)
			if !isError {
				continue
			}

			if err != nil {
				log.Error(err, "Get instance error")
			}
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
	}
}

func (c *constructor) Check() bool {
	return c.kindCheck
}

func (c *constructor) IsFunc() *constructor {
	c.kindCheck = c.t.Kind() == reflect.Func
	return c
}

func (c *constructor) IsReturnMatch(returnType reflect.Type) *constructor {
	if c.t.NumOut() == 0 {
		c.returnCheck = false
		return c
	}

	if c.t.Out(0) != returnType {
		c.returnCheck = false
		return c
	}

	c.returnCheck = true
	return c
}
