package middlewares

import (
	"github.com/lowl11/boost/data/domain"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/io/types"
	"github.com/lowl11/boost/pkg/system/di"
	"time"
)

func Cache(expire time.Duration) domain.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		c := di.Interface[interfaces.Cache]()
		if c == nil {
			return ctx.Next()
		}

		key := "rest" + types.String(ctx.Request().URI().RequestURI())

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
			if ctx.Response().StatusCode() > 299 {
				return ctx.Next()
			}

			if err = c.Set(ctx.Context(), key, types.String(ctx.Response().Body()), expire); err != nil {
				return ctx.Error(err)
			}
		}

		return ctx.Next()
	}
}
