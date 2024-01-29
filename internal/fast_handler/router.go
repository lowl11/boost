package fast_handler

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/types"
	"strings"
	"sync"
)

type router struct {
	routes sync.Map
}

func newRouter() *router {
	return &router{}
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

func (router *router) Register(
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

func (router *router) Get(path string) (*RouteContext, bool) {
	route, ok := router.routes.Load(path)
	if !ok {
		return nil, false
	}

	return route.(*RouteContext), true
}

func (router *router) Search(searchPath string) (*RouteContext, bool) {
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
	if isLastSlash(searchPath) {
		searchPath = removeLast(searchPath)
	}

	// search route...
	router.routes.Range(func(routePath, routeCtx any) bool {
		routePathString := routePath.(string)

		// if routePath ends with '/'
		if isLastSlash(routePathString) {
			routePathString = removeLast(routePathString)
		}

		if searcher := newSearcher(searchPath, routePathString); searcher.Find() {
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
