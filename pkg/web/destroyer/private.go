package destroyer

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
	"github.com/lowl11/boost/pkg/system/types"
)

func (destroyer *Destroyer) runFunc(action types.DestroyFunc) {
	if err := exception.Try(func() error {
		action()
		return nil
	}); err != nil {
		log.Error("Catch panic from destroy action:", err)
	}
}
