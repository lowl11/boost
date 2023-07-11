package context

import (
	"encoding/json"
	"github.com/lowl11/boost/pkg/context"
	"github.com/valyala/fasthttp"
)

func (ctx *Context) Request() *fasthttp.Request {
	return &ctx.inner.Request
}

func (ctx *Context) Status(status int) context.IBoostContext {
	ctx.status = status
	return ctx
}

func (ctx *Context) JSON(body any) error {
	bodyInBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ctx.inner.SetStatusCode(ctx.status)
	ctx.inner.Response.Header.Set("Content-Type", "application/json")
	ctx.inner.SetBody(bodyInBytes)

	return nil
}

func (ctx *Context) String(message string) error {
	ctx.inner.SetStatusCode(ctx.status)
	ctx.inner.Response.Header.Set("Content-Type", "text/plain")
	ctx.inner.SetBody([]byte(message))

	return nil
}
