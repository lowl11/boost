package boost

type group struct {
	router Router
	base   string
}

func newGroup(router Router, base string) *group {
	return &group{
		router: router,
		base:   base,
	}
}

func (group *group) ANY(path string, action HandlerFunc) {
	endpoint := group.base + path
	group.router.ANY(endpoint, action)
}

func (group *group) GET(path string, action HandlerFunc) {
	endpoint := group.base + path
	group.router.GET(endpoint, action)
}

func (group *group) POST(path string, action HandlerFunc) {
	endpoint := group.base + path
	group.router.POST(endpoint, action)
}

func (group *group) PUT(path string, action HandlerFunc) {
	endpoint := group.base + path
	group.router.PUT(endpoint, action)
}

func (group *group) DELETE(path string, action HandlerFunc) {
	endpoint := group.base + path
	group.router.DELETE(endpoint, action)
}
