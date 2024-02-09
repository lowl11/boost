package boost

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/healthcheck"
	"github.com/lowl11/boost/internal/stat"
	"github.com/lowl11/boost/pkg/system/types"
)

func registerStaticEndpoints(app routing, healthcheck *healthcheck.Healthcheck) {
	// register healthcheck endpoint
	app.GET("/health", staticEndpointHealthcheck(healthcheck))

	// register ping/pong endpoint
	app.GET("/ping", staticEndpointPingPong())

	// register stat endpoint
	app.GET("/stat", staticEndpointStat(healthcheck))
}

func staticEndpointHealthcheck(healthcheck *healthcheck.Healthcheck) types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		if err := healthcheck.Trigger(); err != nil {
			return ctx.Error(err)
		}

		return ctx.Ok("OK")
	}
}

func staticEndpointPingPong() types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		return ctx.Ok("pong")
	}
}

func staticEndpointStat(healthcheck *healthcheck.Healthcheck) types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		return ctx.Ok(stat.Format(healthcheck))
	}
}
