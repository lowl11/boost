package context

import (
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/boost/pkg/boost_context"
	"github.com/valyala/fasthttp"
	"net/url"
	"strings"
)

func (ctx *Context) Request() *fasthttp.Request {
	return &ctx.inner.Request
}

func (ctx *Context) Param(name string) string {
	panic("implement me")
}

func (ctx *Context) QueryParam(name string) string {
	uri, err := url.ParseRequestURI(string(ctx.inner.RequestURI()))
	if err != nil {
		return ""
	}

	queryParam, ok := uri.Query()[name]
	if !ok || len(queryParam) == 0 {
		return ""
	}

	return queryParam[0]
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
