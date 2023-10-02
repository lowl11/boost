package destroyer

import (
	"github.com/lowl11/boost/pkg/system/types"
)

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
	}
}
