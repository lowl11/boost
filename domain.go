package boost

import "github.com/lowl11/boost/pkg/interfaces"

type (
	HandlerFunc    = func(ctx Context) error
	MiddlewareFunc = func(ctx Context) error
	Context        = interfaces.Context
	Error          = interfaces.Error
	Route          = interfaces.Route
)

type routing interface {
	ANY(path string, action HandlerFunc) Route
	GET(path string, action HandlerFunc) Route
	POST(path string, action HandlerFunc) Route
	PUT(path string, action HandlerFunc) Route
	DELETE(path string, action HandlerFunc) Route
}

type Router interface {
	routing

	Group(base string) Group
}

type Group interface {
	routing
}
