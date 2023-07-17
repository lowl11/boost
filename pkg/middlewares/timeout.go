package middlewares

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/types"
)

func Timeout() types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		if err := ctx.Next(); err != nil {
			return err
		}

		return nil
	}
}
