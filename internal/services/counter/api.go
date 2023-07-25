package counter

func (counter *Counter) GetRoutes() int {
	return int(counter.routes.Load())
}

func (counter *Counter) GetGroups() int {
	return int(counter.groups.Load())
}

func (counter *Counter) GetMiddlewares() int {
	return int(counter.middlewares.Load())
}

func (counter *Counter) GetGlobalMiddlewares() int {
	return int(counter.globalMiddlewares.Load())
}

func (counter *Counter) GetGroupMiddlewares() int {
	return int(counter.groupMiddlewares.Load())
}

func (counter *Counter) GetRouteMiddlewares() int {
	return int(counter.routeMiddlewares.Load())
}

func (counter *Counter) Route() {
	counter.routes.Add(1)
}

func (counter *Counter) Group() {
	counter.groups.Add(1)
}

func (counter *Counter) GlobalMiddleware() {
	counter.middlewares.Add(1)
	counter.globalMiddlewares.Add(1)
}

func (counter *Counter) GroupMiddleware() {
	counter.middlewares.Add(1)
	counter.groupMiddlewares.Add(1)
}

func (counter *Counter) RouteMiddleware() {
	counter.middlewares.Add(1)
	counter.routeMiddlewares.Add(1)
}
