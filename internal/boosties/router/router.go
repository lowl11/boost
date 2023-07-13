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
}

func (route *RouteContext) Use(middlewareFunc ...func(ctx interfaces.Context) error) {
	if len(middlewareFunc) == 0 {
		return
	}

	//
}
