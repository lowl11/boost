package context

import (
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

type Context struct {
	inner *fasthttp.RequestCtx

	status       int
	keyContainer sync.Map
	params       map[string]string
}

func New(inner *fasthttp.RequestCtx) *Context {
	return &Context{
		inner:  inner,
		status: http.StatusOK,
		params: make(map[string]string),
	}
}
