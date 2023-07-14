package boost

import (
	"github.com/lowl11/boost/internal/boosties/printer"
	"github.com/lowl11/boost/pkg/types"
	"github.com/lowl11/lazylog/log"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

const (
	methodAny  = "ANY"
	emptyGroup = ""
)

func (app *App) Run(port string) {
	printer.PrintGreeting()
	log.Fatal(app.handler.Run(port))
}

func (app *App) Destroy(destroyFunc func()) {
	app.destroyer.AddFunction(destroyFunc)
}

func (app *App) ANY(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(methodAny, path, action, emptyGroup)
}

func (app *App) GET(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodGet, path, action, emptyGroup)
}

func (app *App) POST(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPost, path, action, emptyGroup)
}

func (app *App) PUT(path string, action HandlerFunc) Route {
	return app.handler.RegisterRoute(http.MethodPut, path, action, emptyGroup)
}

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

	app.handler.RegisterGlobalMiddlewares(middlewares...)
}

func (app *App) useGroup(groupID uuid.UUID, middlewareFunc ...MiddlewareFunc) {
	if len(middlewareFunc) == 0 {
		return
	}

	middlewares := make([]types.MiddlewareFunc, 0, len(middlewareFunc))

	for _, mFunc := range middlewareFunc {
		middlewares = append(middlewares, mFunc)
	}

	app.handler.RegisterGroupMiddlewares(groupID, middlewares...)
}
