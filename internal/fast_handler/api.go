package fast_handler

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/types"
	uuid "github.com/satori/go.uuid"
)

func (handler *Handler) Run(port string) error {
	// prepare server
	handler.server.Handler = handler.handler

	// run server
	return handler.server.ListenAndServe(port)
}

func (handler *Handler) RegisterRoute(method, path string, action types.HandlerFunc, groupID string) interfaces.Route {
	if action == nil {
		panic("route action is NULL")
	}

	return handler.router.Register(method, path, action, groupID)
}

func (handler *Handler) RegisterGlobalMiddlewares(middlewareFunc ...types.MiddlewareFunc) {
	middlewareHandlers := make([]types.HandlerFunc, 0, len(middlewareFunc))
	for _, middleware := range middlewareFunc {
		middlewareHandlers = append(middlewareHandlers, types.HandlerFunc(middleware))
	}
	handler.globalMiddlewares = append(handler.globalMiddlewares, middlewareHandlers...)
}

func (handler *Handler) RegisterGroupMiddlewares(groupID uuid.UUID, middlewareFunc ...types.MiddlewareFunc) {
	middlewareHandlers := make([]types.HandlerFunc, 0, len(middlewareFunc))
	for _, middleware := range middlewareFunc {
		middlewareHandlers = append(middlewareHandlers, types.HandlerFunc(middleware))
	}
	handler.groupMiddlewares[groupID.String()] = middlewareHandlers
}
