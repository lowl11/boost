package boost

import (
	"github.com/lowl11/boost/internal/boosties/printer"
	"github.com/lowl11/boost/pkg/types"
	"github.com/lowl11/boostcron"
	"github.com/lowl11/lazylog/log"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

const (
	methodAny  = "ANY"
	emptyGroup = ""
)

// Run starts listening TCP with given port
func (app *App) Run(port string) {
	// register static endpoints
	registerStaticEndpoints(app, app.healthcheck)

	// print greeting text
	printer.PrintGreeting()

	// run cron
	if app.cron != nil {
		app.cron.RunAsync()
	}

	// run server app
	log.Fatal(app.handler.Run(port))
}

// Destroy adds function which will be called in shutdown
func (app *App) Destroy(destroyFunc types.DestroyFunc) {
	if destroyFunc == nil {
		return
	}

	app.destroyer.AddFunction(destroyFunc)
}

// Healthcheck add new application service to healthcheck
func (app *App) Healthcheck(name, url string) {
	app.healthcheck.Register(name, url)
}

// ANY add new route to App with method ANY.
// Note: ANY will receive any type of HTTP method
func (app *App) ANY(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(methodAny, path, action, emptyGroup)
}

// GET add new route to App with method GET
func (app *App) GET(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodGet, path, action, emptyGroup)
}

// POST add new route to App with method POST
func (app *App) POST(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPost, path, action, emptyGroup)
}

// PUT add new route to App with method PUT
func (app *App) PUT(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPut, path, action, emptyGroup)
}

// DELETE add new route to App with method DELETE
func (app *App) DELETE(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodDelete, path, action, emptyGroup)
}

func (app *App) groupANY(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(methodAny, path, action, groupID)
}

func (app *App) groupGET(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodGet, path, action, groupID)
}

func (app *App) groupPOST(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodPost, path, action, groupID)
}

func (app *App) groupPUT(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodPut, path, action, groupID)
}

func (app *App) groupDELETE(path string, action HandlerFunc, groupID string) Route {
	return app.handler.RegisterRoute(http.MethodDelete, path, action, groupID)
}

// Group creates new group for endpoints with base url/endpoint
func (app *App) Group(base string) Group {
	return newGroup(app, base)
}

// CronApp returns boost cron application
func (app *App) CronApp() boostcron.CronRouter {
	if app.cron == nil {
		app.cron = boostcron.New(app.config.CronConfig)
	}

	return app.cron
}

// Use add usable middleware to global App.
// Note: this method adds middleware function to every endpoint
func (app *App) Use(middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	middlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		if mFunc == nil {
			continue
		}

		middlewares = append(middlewares, mFunc)
	}

	app.handler.RegisterGlobalMiddlewares(middlewares...)
}

func (app *App) useGroup(groupID uuid.UUID, middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	middlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		if mFunc == nil {
			continue
		}

		middlewares = append(middlewares, mFunc)
	}

	app.handler.RegisterGroupMiddlewares(groupID, middlewares...)
}
