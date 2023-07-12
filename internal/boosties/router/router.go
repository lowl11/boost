package router

import (
	"github.com/lowl11/boost/pkg/boost_handler"
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
	Action boost_handler.HandlerFunc

	WaitParam bool
	Params    map[string]string
}
