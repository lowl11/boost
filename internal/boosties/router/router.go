package router

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/boost/pkg/types"
	"sync"
)

type Router struct {
	routes sync.Map
}

func New() *Router {
	return &Router{}
}

type RouteContext struct {
	Path   string
	Method string
	Action types.HandlerFunc

	WaitParam bool
	Params    map[string]string

	GroupID     string
	Middlewares []types.HandlerFunc
}

func (route *RouteContext) Use(middlewares ...func(ctx interfaces.Context) error) {
	if len(middlewares) == 0 {
		return
	}

	middlewareHandlers := make([]types.HandlerFunc, 0, len(middlewares))
	for _, middleware := range middlewares {
		if middleware == nil {
			continue
		}

		middlewareHandlers = append(middlewareHandlers, middleware)
	}

	route.Middlewares = middlewareHandlers
}
