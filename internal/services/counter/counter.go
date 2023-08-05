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
