package di_container

import (
	"github.com/lowl11/boost/data/enums/di_modes"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/io/flex"
	"github.com/lowl11/boost/pkg/io/types"
	"reflect"
	"sync"
)

type serviceInfo struct {
	constructor any
	mode        int
	instance    any
	tq          *tQueue
}

type container struct {
	services                    map[reflect.Type]*serviceInfo
	controllerInterfaceType     reflect.Type
	registerEndpointsMethodName string
	appType                     reflect.Type

	isChecked bool

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

func (c *container) SetControllerInterface(controllerInterfaceType reflect.Type) {
	c.controllerInterfaceType = types.UnwrapType(controllerInterfaceType)
	c.registerEndpointsMethodName = c.controllerInterfaceType.Method(0).Name
}

func (c *container) SetAppType(appType reflect.Type) {
	c.appType = types.UnwrapType(appType)
}

func (c *container) Register(t reflect.Type, constructor any, mode int, dependencies ...any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	t = types.UnwrapType(t)

	// check mode
	if mode < 1 || mode > 3 {
		return
	}

	// check constructor
	newConstructor(constructor).
		IsFunc().
		IsReturnMatch(t)

	tq := newTQueue()
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

func (c *container) RegisterImplementation(impl any) {
	c.services[types.UnwrapType(reflect.TypeOf(impl))] = &serviceInfo{
		mode:     di_modes.Singleton,
		instance: impl,
	}
}

func (c *container) MapControllers(constructors ...any) {
	// create singleton no dep service instances
	for serviceType, service := range c.services {
		// instance already exist or service mode is not singleton
		if service.instance != nil || service.mode != di_modes.Singleton {
			continue
		}

		// skip, because there is non-empty arguments
		f, _ := flex.Func(service.constructor)
		if f.ArgumentsCount() > 0 {
			continue
		}

		// create instance
		serviceInstance := call(service.tq, service.constructor, c.services)
		if serviceInstance == nil {
			continue // todo: check
		}

		// store instance
		c.services[serviceType].instance = serviceInstance
	}

	// create singleton services more than 0 arguments
	for serviceType, service := range c.services {
		if service.instance != nil || service.mode != di_modes.Singleton {
			continue
		}

		serviceInstance := call(service.tq, service.constructor, c.services)
		if serviceInstance == nil {
			continue // todo: check
		}

		c.services[serviceType].instance = serviceInstance
	}

	// register controller types
	controllerTypes := make([]reflect.Type, 0, len(constructors))
	for _, controllerConstructor := range constructors {
		flexConstructor, _ := flex.Func(controllerConstructor)
		if flexConstructor.ReturnsCount() != 1 {
			panic("Controller Constructor has no return value or too much values: " + reflect.TypeOf(controllerConstructor).String())
		}

		controllerType := flexConstructor.Returns()[0]

		if !controllerType.Implements(c.controllerInterfaceType) {
			panic("Given controller does not implements boost.Controller interface: " + controllerType.String())
		}

		c.Register(controllerType, controllerConstructor, di_modes.Singleton)
		controllerTypes = append(controllerTypes, controllerType)
	}

	// init controllers
	for _, controllerType := range controllerTypes {
		controllerObject := c.Get(controllerType)
		registerMethod, _ := controllerType.MethodByName(c.registerEndpointsMethodName)
		registerMethod.Func.Call([]reflect.Value{
			reflect.ValueOf(controllerObject),
			reflect.ValueOf(c.Get(c.appType)),
		})
	}
}

func (c *container) Get(t reflect.Type) any {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	t = types.UnwrapType(t)

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

	if c.isChecked {
		return nil
	}

	for sType, sInfo := range c.services {
		if sInfo.instance != nil {
			continue
		}

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
					panic("No set primitive dependency at registry: " + dep.String() + " for " + sType.String())
				}
			}
		}
	}

	c.isChecked = true

	return nil
}
