package boost

import (
	"github.com/lowl11/boost/pkg/interfaces"
	uuid "github.com/satori/go.uuid"
)

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

type groupRouting interface {
	groupANY(path string, action HandlerFunc, groupID string) Route
	groupGET(path string, action HandlerFunc, groupID string) Route
	groupPOST(path string, action HandlerFunc, groupID string) Route
	groupPUT(path string, action HandlerFunc, groupID string) Route
	groupDELETE(path string, action HandlerFunc, groupID string) Route
}

type Router interface {
	routing

	Group(base string) Group
	useGroup(groupID uuid.UUID, middlewares ...MiddlewareFunc)
}

type groupRouter interface {
	groupRouting

	useGroup(groupID uuid.UUID, middlewares ...MiddlewareFunc)
}

type Group interface {
	routing

	Use(middlewares ...MiddlewareFunc)
}
