package di_container

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"reflect"
)

type serviceInfo struct {
	constructor any
	mode        int
	instance    any
}

type container struct {
	services         map[reflect.Type]*serviceInfo
	sessionInstances map[string][]*serviceInfo
}

func newContainer() *container {
	return &container{
		services:         make(map[reflect.Type]*serviceInfo, 10),
		sessionInstances: make(map[string][]*serviceInfo),
	}
}

var instance *container

func Get() interfaces.DependencyContainer {
	if instance != nil {
		return instance
	}

	instance = newContainer()
	return instance
}

func (c *container) Register(t reflect.Type, constructor any, mode int) {
	t = type_helper.UnwrapType(t)

	// check mode
	if mode < 1 || mode > 3 {
		return
	}

	// check constructor
	if !newConstructor(constructor).
		IsFunc().
		IsReturnMatch(t).
		Check() {
		return
	}

	// add service
	c.services[t] = &serviceInfo{
		constructor: constructor,
		mode:        mode,
	}
}

func (c *container) Get(t reflect.Type, params ...any) any {
	t = type_helper.UnwrapType(t)

	service, ok := c.services[t]
	if !ok {
		return nil
	}

	switch service.mode {
	case di_modes.Transient:
		// call every time new instance
		return call(service.constructor, c.services)
	case di_modes.Scoped:
		// call one instance for request
		panic("not implemented")
	case di_modes.Singleton:
		// call one instance for all app lifetime
		if service.instance != nil {
			return service.instance
		}

		// return the instance & cache it
		inst := call(service.constructor, c.services)
		service.instance = inst
		return inst
	}

	return nil
}
