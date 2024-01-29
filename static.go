package boost

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/healthcheck"
	"github.com/lowl11/boost/pkg/system/types"
)

func registerStaticEndpoints(app routing, healthcheck *healthcheck.Healthcheck) {
	// register healthcheck endpoints
	app.GET("/health", staticEndpointHealthcheck(healthcheck))

	// register ping/pong endpoints
	app.GET("/ping", staticEndpointPingPong())
}

func staticEndpointHealthcheck(healthcheck *healthcheck.Healthcheck) types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		if err := healthcheck.Trigger(); err != nil {
			return ctx.Error(err)
		}

		return ctx.String("OK")
	}
}

func staticEndpointPingPong() types.HandlerFunc {
	return func(ctx interfaces.Context) error {
		return ctx.String("pong")
	}
}
