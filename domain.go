package boost

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/pkg/interfaces"
)

type (
	// HandlerFunc is REST handler for actions
	HandlerFunc = func(ctx Context) error

	// MiddlewareFunc is type of middleware functions
	MiddlewareFunc = func(ctx Context) error

	// Context is interface for HandlerFunc context argument
	Context = interfaces.Context

	// Error is custom error for Boost
	Error = interfaces.Error

	// Route is interface which will return after adding new route
	Route = interfaces.Route

	// CacheRepository is interface for using cache
	CacheRepository = interfaces.CacheRepository
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

// Group is interface which will return after creating new group
type Group interface {
	routing

	Use(middlewares ...MiddlewareFunc)
}
