package printer

import (
	"fmt"
	"github.com/lowl11/boost/internal/boosties/static_controller"
	"github.com/lowl11/boost/internal/services/counter"
)

const (
	greetingText = `

  | |__   ___   ___  ___| |_ 
  | '_ \ / _ \ / _ \/ __| __|
  | |_) | (_) | (_) \__ \ |_ 
  |_.__/ \___/ \___/|___/\__|
  Minimalist Go framework based on FastHTTP
  https://github.com/lowl11/boost
--------------------------------------------
Routes: %d
Groups: %d
--------------------------------------------
`
)

func PrintGreeting(counter *counter.Counter) {
	fmt.Printf(
		greetingText,
		counter.GetRoutes()-static_controller.RouteCount,
		counter.GetGroups(),
		//counter.GetMiddlewares(),
		//counter.GetGlobalMiddlewares(),
		//counter.GetGroupMiddlewares(),
		//counter.GetRouteMiddlewares(),
	)
}
