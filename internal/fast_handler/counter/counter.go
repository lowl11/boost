package counter

import "sync/atomic"

type Counter struct {
	routes atomic.Int32
	groups atomic.Int32

	middlewares       atomic.Int32
	globalMiddlewares atomic.Int32
	groupMiddlewares  atomic.Int32
	routeMiddlewares  atomic.Int32

	cronActions atomic.Int32

	listenerBind atomic.Int32
}

func New() *Counter {
	return &Counter{}
}

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

func (counter *Counter) GetCronActions() int {
	return int(counter.cronActions.Load())
}

func (counter *Counter) GetListenerBind() int {
	return int(counter.listenerBind.Load())
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

func (counter *Counter) CronAction() {
	counter.cronActions.Add(1)
}

func (counter *Counter) ListenerBind(value int) {
	counter.listenerBind.Add(int32(value))
}
