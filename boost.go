package boost

import (
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/lazyconfig/config/config_internal"
	"github.com/lowl11/lazylog/log/log_internal"
)

type App struct {
	handler *fast_handler.Handler
}

func New() *App {
	log_internal.Init(log_internal.LogConfig{})
	config_internal.Init(config_internal.Config{})

	return &App{
		handler: fast_handler.New(),
	}
}

type (
	HandlerFunc    = func(ctx Context) error
	MiddlewareFunc = func(ctx Context) error
	Context        = interfaces.Context
	Error          = interfaces.Error
	Route          = interfaces.Route
)

type routing interface {
	ANY(path string, action HandlerFunc) Route
	GET(path string, action HandlerFunc) Route
	POST(path string, action HandlerFunc) Route
	PUT(path string, action HandlerFunc) Route
	DELETE(path string, action HandlerFunc) Route
}

type Router interface {
	routing

	Group(base string) Group
}

type Group interface {
	routing
}
