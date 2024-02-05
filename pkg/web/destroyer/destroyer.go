package destroyer

import (
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
