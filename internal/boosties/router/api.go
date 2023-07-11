package router

import (
	"github.com/lowl11/boost/pkg/boost_handler"
)

func (router *Router) Add(method, path string, action boost_handler.HandlerFunc) *Router {
	// skip add route if already exist
	if _, exist := router.routes[path]; exist {
		return router
	}

	router.routes[path] = RouteContext{
		Path:   path,
		Method: method,
		Action: action,
	}

	return router
}

func (router *Router) Map() map[string]RouteContext {
	return router.routes
}
