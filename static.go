package boost

import (
	"github.com/lowl11/boost/internal/boosties/static_controller"
	"github.com/lowl11/boost/internal/services/boost/healthcheck"
)

func registerStaticEndpoints(app routing, healthcheck *healthcheck.Healthcheck) {
	// register healthcheck endpoints
	app.GET("/health", static_controller.Healthcheck(healthcheck))

	// register ping/pong endpoints
	app.GET("/ping", static_controller.Ping())
}
