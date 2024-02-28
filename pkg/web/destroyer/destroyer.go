package destroyer

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
	"github.com/lowl11/boost/pkg/system/types"
	"sync"
)

type Destroyer struct {
	functions []types.DestroyFunc
	mutex     sync.Mutex
}

var instance *Destroyer

func Get() *Destroyer {
	if instance != nil {
		return instance
	}

	instance = &Destroyer{
		functions: make([]types.DestroyFunc, 0),
	}
	return instance
}

func (destroyer *Destroyer) AddFunction(destroyFunc types.DestroyFunc) *Destroyer {
	destroyer.mutex.Lock()
	defer destroyer.mutex.Unlock()

	destroyer.functions = append(destroyer.functions, destroyFunc)
	return destroyer
}

func (destroyer *Destroyer) Destroy() {
	destroyer.mutex.Lock()
	defer destroyer.mutex.Unlock()

	for _, destroyFunc := range destroyer.functions {
		destroyer.runFunc(destroyFunc)
		if err := exception.Try(func() error {
			destroyer.runFunc(destroyFunc)
			return nil
		}); err != nil {
			log.Error("Catch panic from destroy action:", err)
		}
	}
}

func (destroyer *Destroyer) runFunc(action types.DestroyFunc) {
	if err := exception.Try(func() error {
		action()
		return nil
	}); err != nil {
		log.Error("Catch panic from destroy action:", err)
	}
}
