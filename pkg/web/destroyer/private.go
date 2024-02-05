package destroyer

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
	"github.com/lowl11/boost/pkg/system/types"
)

func (destroyer *Destroyer) runFunc(action types.DestroyFunc) {
	defer func() {
		if err := exception.CatchPanic(recover()); err != nil {
			log.Error("Catch panic from destroy action:", err)
		}
	}()

	action()
}
