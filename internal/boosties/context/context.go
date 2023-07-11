package context

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

type Context struct {
	inner *fasthttp.RequestCtx

	status int
}

func New(inner *fasthttp.RequestCtx) *Context {
	return &Context{
		inner:  inner,
		status: http.StatusOK,
	}
}
