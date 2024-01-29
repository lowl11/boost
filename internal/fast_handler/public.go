package fast_handler

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/services/counter"
	types2 "github.com/lowl11/boost/pkg/system/types"
	"net"
)

func (handler *Handler) Run(port string) error {
	// prepare server
	handler.server.Handler = handler.handler

	// define listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// check & update CORS
	handler.tryUpdateCORS()

	// run server
	return handler.server.Serve(listener)
}

func (handler *Handler) RegisterRoute(method, path string, action types2.HandlerFunc, groupID string) interfaces.Route {
	if action == nil {
		panic("route action is NULL")
	}

	handler.counter.Route()
	return handler.router.Register(method, path, action, groupID)
}

func (handler *Handler) RegisterGroup() {
	handler.counter.Group()
}

func (handler *Handler) RegisterGlobalMiddlewares(middlewareFunc ...types2.MiddlewareFunc) {
	middlewareHandlers := make([]types2.HandlerFunc, 0, len(middlewareFunc))
	for _, middleware := range middlewareFunc {
		if middleware == nil {
			continue
		}

		handler.counter.GlobalMiddleware()
		middlewareHandlers = append(middlewareHandlers, types2.HandlerFunc(middleware))
	}

	handler.globalMiddlewares = append(handler.globalMiddlewares, middlewareHandlers...)
}

func (handler *Handler) RegisterGroupMiddlewares(groupID uuid.UUID, middlewareFunc ...types2.MiddlewareFunc) {
	middlewareHandlers := make([]types2.HandlerFunc, 0, len(middlewareFunc))
	for _, middleware := range middlewareFunc {
		if middleware == nil {
			continue
		}

		handler.counter.GroupMiddleware()
		middlewareHandlers = append(middlewareHandlers, types2.HandlerFunc(middleware))
	}

	handler.groupMiddlewares[groupID.String()] = middlewareHandlers
}

func (handler *Handler) GetCounter() *counter.Counter {
	return handler.counter
}

func (handler *Handler) SetCorsConfig(config CorsConfig) {
	handler.corsConfig = config
}

func (handler *Handler) SetPanicHandler(panicHandler types2.PanicHandler) {
	handler.panicHandler = panicHandler
}
