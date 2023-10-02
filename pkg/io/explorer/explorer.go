package explorer

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"sync"
)

type Explorer struct {
	path string

	threadSafe bool
	mutex      sync.Mutex
}

func New(path string) interfaces.IExplorer {
	return &Explorer{
		path: path,
	}
}
