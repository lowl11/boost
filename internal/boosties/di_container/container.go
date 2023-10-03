package di_container

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/containers/tqueue"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"reflect"
	"sync"
)

type serviceInfo struct {
	constructor any
	mode        int
	instance    any
	tq          *tqueue.Queue
}

type container struct {
	services map[reflect.Type]*serviceInfo

	mutex sync.Mutex
}

func newContainer() *container {
	return &container{
		services: make(map[reflect.Type]*serviceInfo, 10),
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

func (c *container) Register(t reflect.Type, constructor any, mode int, dependencies ...any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	t = type_helper.UnwrapType(t)

	// check mode
	if mode < 1 || mode > 3 {
		return
	}

	// check constructor
	newConstructor(constructor).
		IsFunc().
		IsReturnMatch(t)

	tq := tqueue.New()
	for _, dep := range dependencies {
		tq.Enqueue(dep)
	}

	// add service
	c.services[t] = &serviceInfo{
		constructor: constructor,
		mode:        mode,
		tq:          tq,
	}
}

func (c *container) Get(t reflect.Type, params ...any) any {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	t = type_helper.UnwrapType(t)

	service, ok := c.services[t]
	if !ok {
		return nil
	}

	switch service.mode {
	case di_modes.Transient:
		// call every time new instance
		return call(service.tq, service.constructor, c.services)
	case di_modes.Scoped:
		// call one instance for request
		panic("not implemented")
	case di_modes.Singleton:
		// call one instance for all app lifetime
		if service.instance != nil {
			return service.instance
		}

		// return the instance & cache it
		inst := call(service.tq, service.constructor, c.services)
		service.instance = inst
		return inst
	}

	return nil
}

func (c *container) Check() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for sType, sInfo := range c.services {
		deps := newConstructor(sInfo.constructor).GetDependencies()
		if len(deps) == 0 {
			continue
		}

		tq := sInfo.tq.Copy()
		for _, dep := range deps {
			if dep.Kind() == reflect.Pointer || dep.Kind() == reflect.Struct {
				if _, depOk := c.services[dep]; !depOk {
					panic("The dependency " + dep.String() + " is not registered for " + sType.String())
				}
			} else {
				primitiveValue := tq.Dequeue(dep)
				if primitiveValue == nil {
					panic("No set primitive dependency at registry: " + dep.String())
				}
			}
		}
	}

	return nil
}
