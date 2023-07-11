package fast_handler

import (
	"github.com/lowl11/boost/internal/boosties/router"
	"github.com/valyala/fasthttp"
)

func (handler *Handler) Run(port string) error {
	return fasthttp.ListenAndServe(port, handler.commonHandler)
}

func (handler *Handler) AddRoute(method, path string, action any) {
	h, ok := action.(router.HandlerFunc)
	if !ok {
		return
	}

	handler.router.Add(method, path, h)
}
