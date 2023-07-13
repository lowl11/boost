package boost

import (
	"github.com/lowl11/boost/internal/boosties/printer"
	"github.com/lowl11/boost/pkg/types"
	"github.com/lowl11/lazylog/log"
	"net/http"
)

const (
	methodAny = "ANY"
)

func (app *App) Run(port string) {
	printer.PrintGreeting()
	log.Fatal(app.handler.Run(port))
}

func (app *App) Destroy(destroyFunc func()) {
	app.destroyer.AddFunction(destroyFunc)
}

func (app *App) ANY(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(methodAny, path, action)
}

func (app *App) GET(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodGet, path, action)
}

func (app *App) POST(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPost, path, action)
}

func (app *App) PUT(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPut, path, action)
}

func (app *App) DELETE(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodDelete, path, action)
}

func (app *App) Group(base string) Group {
	return newGroup(app, base)
}

func (app *App) Use(middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	middlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		middlewares = append(middlewares, mFunc)
	}

	app.handler.RegisterMiddleware(middlewares...)
}
