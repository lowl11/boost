package boost

import (
	"github.com/lowl11/boost/internal/boosties/printer"
	"github.com/lowl11/lazylog/log"
	"net/http"
)

const (
	methodAny = "ANY"
)

func (boost *App) Run(port string) {
	printer.PrintGreeting()
	log.Fatal(boost.handler.Run(port))
}

func (boost *App) ANY(path string, action HandlerFunc) {
	boost.handler.RegisterRoute(methodAny, path, action)
}

func (boost *App) GET(path string, action HandlerFunc) {
	boost.handler.RegisterRoute(http.MethodGet, path, action)
}

func (boost *App) POST(path string, action HandlerFunc) {
	boost.handler.RegisterRoute(http.MethodPost, path, action)
}

func (boost *App) PUT(path string, action HandlerFunc) {
	boost.handler.RegisterRoute(http.MethodPut, path, action)
}

func (boost *App) DELETE(path string, action HandlerFunc) {
	boost.handler.RegisterRoute(http.MethodDelete, path, action)
}
