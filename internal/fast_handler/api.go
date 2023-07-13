package fast_handler

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/types"
	"github.com/valyala/fasthttp"
)

func (handler *Handler) Run(port string) error {
	// run server
	return fasthttp.ListenAndServe(port, handler.commonHandler)
}

func (handler *Handler) RegisterRoute(method, path string, action types.HandlerFunc) interfaces.Route {
	return handler.router.Register(method, path, action)
}

func (handler *Handler) RegisterMiddleware(middlewareFunc ...types.MiddlewareFunc) {
	//
}
