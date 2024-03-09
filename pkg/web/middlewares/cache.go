package middlewares

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/pkg/system/types"
	"time"
)

func Cache(expire time.Duration) types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		c := di.Interface[interfaces.Cache]()
		if c == nil {
			return ctx.Next()
		}

		key := "rest" + types.ToString(ctx.Request().URI().RequestURI())

		log.Debug(key)

		content, err := c.Get(ctx.Context(), key)
		if err != nil {
			return ctx.Error(err)
		}

		if content != nil {
			return ctx.
				SetHeader("Content-Type", "application/json").
				Bytes(content)
		}

		if err = ctx.Next(); err != nil {
			return ctx.Error(err)
		}

		if content == nil {
			if err = c.Set(ctx.Context(), key, types.ToString(ctx.Response().Body()), expire); err != nil {
				return ctx.Error(err)
			}
		}

		return nil
	}
}
