package boost

import (
	"github.com/lowl11/boost/internal/fast_handler"
	"github.com/lowl11/boost/pkg/boost_context"
	"github.com/lowl11/boost/pkg/boost_handler"
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
	HandlerFunc = boost_handler.HandlerFunc
	Context     = boost_context.Context
)
