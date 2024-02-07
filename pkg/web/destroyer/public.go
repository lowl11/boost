package destroyer

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
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
		if err := exception.Try(func() error {
			destroyer.runFunc(destroyFunc)
			return nil
		}); err != nil {
			log.Error("Catch panic from destroy action:", err)
		}
	}
}
