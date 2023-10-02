package router

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/helpers/path_helper"
	"github.com/lowl11/boost/internal/services/boost/route_searcher"
	"github.com/lowl11/boost/pkg/system/types"
	"strings"
)

func (router *Router) Register(
	method, path string,
	action types.HandlerFunc,
	groupID string,
) interfaces.Route {
	// if path contains dynamic params
	var waitParam bool
	if strings.Contains(path, ":") {
		waitParam = true
	}

	route := &RouteContext{
		Path:   path,
		Method: method,
		Action: action,

		WaitParam: waitParam,
		GroupID:   groupID,
	}

	// skip add route if already exist
	if _, exist := router.routes.Load(path); exist {
		return route
	}

	// register new route
	router.routes.Store(path, route)

	return route
}

func (router *Router) Get(path string) (*RouteContext, bool) {
	route, ok := router.routes.Load(path)
	if !ok {
		return nil, false
	}

	return route.(*RouteContext), true
}

func (router *Router) Search(searchPath string) (*RouteContext, bool) {
	// try to find by static path
	route, ok := router.Get(searchPath)

	// found by static path
	if ok {
		return route, true
	}

	// not found by static path, may have variable
	var searchRoute *RouteContext
	var found bool

	// if searchPath ends with '/'
	if path_helper.IsLastSlash(searchPath) {
		searchPath = path_helper.RemoveLast(searchPath)
	}

	// search route...
	router.routes.Range(func(routePath, routeCtx any) bool {
		routePathString := routePath.(string)

		// if routePath ends with '/'
		if path_helper.IsLastSlash(routePathString) {
			routePathString = path_helper.RemoveLast(routePathString)
		}

		if searcher := route_searcher.New(searchPath, routePathString); searcher.Find() {
			found = true
			searchRoute = routeCtx.(*RouteContext)
			searchRoute.Params = searcher.Params()
			return false
		}

		// keep going
		return true
	})

	// if route found by search - return it
	if found {
		return searchRoute, true
	}

	// not found case
	return nil, false
}
