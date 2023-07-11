package router

func (router *Router) Add(method, path string, action HandlerFunc) *Router {
	router.routes[path] = routeItem{
		path:   path,
		method: method,
		action: action,
	}

	return router
}
