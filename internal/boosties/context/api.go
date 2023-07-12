package context

import (
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/boost/pkg/boost_context"
	"github.com/valyala/fasthttp"
	"strings"
)

func (ctx *Context) Request() *fasthttp.Request {
	return &ctx.inner.Request
}

func (ctx *Context) Param(name string) string {
	return ctx.params[name]
}

func (ctx *Context) SetParams(params map[string]string) *Context {
	if params == nil {
		return ctx
	}

	ctx.params = params
	return ctx
}

func (ctx *Context) QueryParam(name string) string {
	return string(ctx.inner.URI().QueryArgs().Peek(name))
}

func (ctx *Context) IsWebSocket() bool {
	headerUpgrade := string(ctx.inner.Request.Header.Peek("Upgrade"))
	return strings.EqualFold(headerUpgrade, "websocket")
}

func (ctx *Context) Set(key string, value any) {
	ctx.keyContainer.Store(key, value)
}

func (ctx *Context) Get(key string) any {
	value, ok := ctx.keyContainer.Load(key)
	if !ok {
		return nil
	}

	return value
}

func (ctx *Context) Status(status int) boost_context.Context {
	ctx.status = status
	return ctx
}

func (ctx *Context) Empty() error {
	ctx.inner.SetStatusCode(ctx.status)
	ctx.inner.Response.Header.Set("Content-Type", "application/json")
	ctx.inner.SetBody(nil)

	return nil
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

func (ctx *Context) XML(body any) error {
	bodyInBytes, err := xml.Marshal(body)
	if err != nil {
		return err
	}

	ctx.inner.SetStatusCode(ctx.status)
	ctx.inner.Response.Header.Set("Content-Type", "application/xml")
	ctx.inner.SetBody(bodyInBytes)

	return nil
}

func (ctx *Context) String(message string) error {
	ctx.inner.SetStatusCode(ctx.status)
	ctx.inner.Response.Header.Set("Content-Type", "text/plain")
	ctx.inner.SetBody([]byte(message))

	return nil
}
