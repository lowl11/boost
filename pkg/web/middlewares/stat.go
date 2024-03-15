package middlewares

import (
	"github.com/lowl11/boost/data/domain"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/healthcheck"
	"github.com/lowl11/boost/internal/stat"
	"github.com/lowl11/boost/pkg/system/di"
)

func Stat() domain.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		return ctx.Ok(stat.Format(di.Get[healthcheck.Healthcheck]()))
	}
}
