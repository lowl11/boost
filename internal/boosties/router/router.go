package router

import (
	"github.com/lowl11/boost/pkg/context"
)

type Router struct {
	routes map[string]routeItem
}

func New() *Router {
	return &Router{
		routes: make(map[string]routeItem),
	}
}

type HandlerFunc func(ctx context.IBoostContext) error

type routeItem struct {
	path   string
	method string
	action HandlerFunc
}
