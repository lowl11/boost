package boost

import (
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/lazylog/log/log_internal"
)

type App struct {
	handler *fast_handler.Handler
}

func New() *App {
	log_internal.Init(log_internal.LogConfig{})

	return &App{
		handler: fast_handler.New(),
	}
}

type (
	HandlerFunc = func(ctx Context) error
	Context     = interfaces.Context
	Error       = interfaces.Error
)

type routing interface {
	ANY(path string, action HandlerFunc)
	GET(path string, action HandlerFunc)
	POST(path string, action HandlerFunc)
	PUT(path string, action HandlerFunc)
	DELETE(path string, action HandlerFunc)
}

type Router interface {
	routing

	Group(base string) Group
}

type Group interface {
	routing
}
