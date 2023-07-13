package destroyer

import "sync"

type Destroyer struct {
	functions []func()
	mutex     sync.Mutex
}

func New() *Destroyer {
	return &Destroyer{
		functions: make([]func(), 0),
	}
}
