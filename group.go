package boost

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/pkg/web/socket"
)

type group struct {
	id     uuid.UUID
	router groupRouter
	base   string
}

func newGroup(router groupRouter, base string) *group {
	return &group{
		id:     uuid.New(),
		router: router,
		base:   base,
	}
}

func (group *group) ANY(path string, action HandlerFunc) Route {
	endpoint := group.base + path
	return group.router.groupANY(endpoint, action, group.id.String())
}

func (group *group) GET(path string, action HandlerFunc) Route {
	endpoint := group.base + path
	return group.router.groupGET(endpoint, action, group.id.String())
}

func (group *group) POST(path string, action HandlerFunc) Route {
	endpoint := group.base + path
	return group.router.groupPOST(endpoint, action, group.id.String())
}

func (group *group) PUT(path string, action HandlerFunc) Route {
	endpoint := group.base + path
	return group.router.groupPUT(endpoint, action, group.id.String())
}

func (group *group) DELETE(path string, action HandlerFunc) Route {
	endpoint := group.base + path
	return group.router.groupDELETE(endpoint, action, group.id.String())
}

func (group *group) Use(middlewareFunc ...MiddlewareFunc) {
	group.router.useGroup(group.id, middlewareFunc...)
}

func (group *group) Websocket(path string, handler *socket.Handler) {
	websocketHandler(group, path, handler)
}
