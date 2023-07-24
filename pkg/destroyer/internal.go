package destroyer

import (
	"errors"
	"github.com/lowl11/boost/pkg/types"
	"github.com/lowl11/lazylog/log"
)

func (destroyer *Destroyer) runFunc(action types.DestroyFunc) {
	defer func() {
		if value := recover(); value != nil {
			var err error

			if _, ok := value.(string); ok {
				err = errors.New(value.(string))
			} else if _, ok = value.(error); ok {
				err = value.(error)
			}

			log.Error(err, "Catch panic from destroy action")
		}
	}()

	action()
}
