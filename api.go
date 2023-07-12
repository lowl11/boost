package boost

import (
	"github.com/lowl11/boost/internal/boosties/printer"
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

func (app *App) ANY(path string, action HandlerFunc) {
	app.handler.RegisterRoute(methodAny, path, action)
}

func (app *App) GET(path string, action HandlerFunc) {
	app.handler.RegisterRoute(http.MethodGet, path, action)
}

func (app *App) POST(path string, action HandlerFunc) {
	app.handler.RegisterRoute(http.MethodPost, path, action)
}

func (app *App) PUT(path string, action HandlerFunc) {
	app.handler.RegisterRoute(http.MethodPut, path, action)
}

func (app *App) DELETE(path string, action HandlerFunc) {
	app.handler.RegisterRoute(http.MethodDelete, path, action)
}

func (app *App) Group(base string) Group {
	return nil
}
