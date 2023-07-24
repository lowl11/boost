package destroyer

import (
	"github.com/lowl11/boost/pkg/types"
	"sync"
)

type Destroyer struct {
	functions []types.DestroyFunc
	mutex     sync.Mutex
}

func New() *Destroyer {
	return &Destroyer{
		functions: make([]types.DestroyFunc, 0),
	}
}
