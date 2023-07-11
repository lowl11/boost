package router

import (
	"github.com/lowl11/boost/pkg/boost_handler"
)

type Router struct {
	routes map[string]RouteContext
}

func New() *Router {
	return &Router{
		routes: make(map[string]RouteContext),
	}
}

type RouteContext struct {
	Path   string
	Method string
	Action boost_handler.HandlerFunc
}
