package context

import (
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/boost/internal/helpers/fast_helper"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/boost_context"
	"github.com/lowl11/boost/pkg/content_types"
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
	return type_helper.BytesToString(ctx.inner.URI().QueryArgs().Peek(name))
}

func (ctx *Context) Header(name string) string {
	return type_helper.BytesToString(ctx.inner.Request.Header.Peek(name))
}

func (ctx *Context) Headers() map[string]string {
	headers := make(map[string]string)

	ctx.inner.Request.Header.VisitAll(func(key, value []byte) {
		headers[type_helper.BytesToString(key)] = type_helper.BytesToString(value)
	})

	return headers
}

func (ctx *Context) Cookie(name string) string {
	return type_helper.BytesToString(ctx.inner.Request.Header.Cookie(name))
}

func (ctx *Context) Cookies() map[string]string {
	cookies := make(map[string]string)

	ctx.inner.Request.Header.VisitAllCookie(func(key, value []byte) {
		cookies[type_helper.BytesToString(key)] = type_helper.BytesToString(value)
	})

	return cookies
}

func (ctx *Context) Body() []byte {
	return ctx.inner.Request.Body()
}

func (ctx *Context) Parse(object any) error {
	contentType := ctx.Header("Content-Type")

	switch contentType {
	case content_types.JSON:
		return json.Unmarshal(ctx.Body(), &object)
	case content_types.XML:
		return xml.Unmarshal(ctx.Body(), &object)
	}

	return ErrorUnknownContentType(contentType)
}

func (ctx *Context) IsWebSocket() bool {
	headerUpgrade := type_helper.BytesToString(ctx.inner.Request.Header.Peek("Upgrade"))
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
	fast_helper.Write(ctx.inner, content_types.Text, ctx.status, nil)

	return nil
}

func (ctx *Context) String(message string) error {
	fast_helper.Write(ctx.inner, content_types.Text, ctx.status, type_helper.StringToBytes(message))

	return nil
}

func (ctx *Context) Bytes(body []byte) error {
	fast_helper.Write(ctx.inner, content_types.Bytes, ctx.status, body)

	return nil
}

func (ctx *Context) JSON(body any) error {
	bodyInBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fast_helper.Write(ctx.inner, content_types.JSON, ctx.status, bodyInBytes)

	return nil
}

func (ctx *Context) XML(body any) error {
	bodyInBytes, err := xml.Marshal(body)
	if err != nil {
		return err
	}

	fast_helper.Write(ctx.inner, content_types.XML, ctx.status, bodyInBytes)

	return nil
}
