package middlewares

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/healthcheck"
	"github.com/lowl11/boost/internal/stat"
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/pkg/system/types"
)

func Stat() types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		return ctx.Ok(stat.Format(di.Get[healthcheck.Healthcheck]()))
	}
}
