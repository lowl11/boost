package static_controller

import (
	"github.com/lowl11/boost/internal/services/healthcheck"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/system/types"
)

const (
	RouteCount = 2
)

func Healthcheck(healthcheck *healthcheck.Healthcheck) types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		if err := healthcheck.Trigger(); err != nil {
			return ctx.Error(err)
		}

		return ctx.String("OK")
	}
}

func Ping() types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		return ctx.String("pong")
	}
}
