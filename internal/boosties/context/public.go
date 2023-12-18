package context

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/boost/data/enums/content_types"
	"github.com/lowl11/boost/data/enums/headers"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/log"
	"github.com/valyala/fasthttp"
	"io"
	"reflect"
	"strings"
)

func (ctx *Context) Request() *fasthttp.Request {
	return &ctx.inner.Request
}

func (ctx *Context) Response() *fasthttp.Response {
	return &ctx.inner.Response
}

func (ctx *Context) Method() string {
	return type_helper.BytesToString(ctx.inner.Method())
}

func (ctx *Context) Scheme() string {
	if ctx.IsTLS() {
		return "https"
	}

	if scheme := ctx.Header(headers.HeaderXForwardedProto); scheme != "" {
		return scheme
	}

	if scheme := ctx.Header(headers.HeaderXForwardedProtocol); scheme != "" {
		return scheme
	}

	if ssl := ctx.Header(headers.HeaderXForwardedSSL); ssl == "on" {
		return "https"
	}

	if scheme := ctx.Header(headers.HeaderXUrlScheme); scheme != "" {
		return scheme
	}

	return "http"
}

func (ctx *Context) Authorization() string {
	_, after, found := strings.Cut(ctx.Header(headers.HeaderAuthorization), " ")
	if !found {
		return ""
	}

	return after
}

func (ctx *Context) Param(name string) interfaces.Param {
	return NewParam(ctx.params[name])
}

func (ctx *Context) Params() map[string]interfaces.Param {
	params := make(map[string]interfaces.Param)
	for key, value := range ctx.params {
		params[key] = NewParam(value)
	}

	return params
}

func (ctx *Context) SetParams(params map[string]string) *Context {
	if params == nil {
		return ctx
	}

	ctx.params = params
	return ctx
}

func (ctx *Context) QueryParam(name string) interfaces.Param {
	return NewParam(type_helper.BytesToString(ctx.inner.URI().QueryArgs().Peek(name)))
}

func (ctx *Context) QueryParams() map[string]interfaces.Param {
	params := make(map[string]interfaces.Param)
	ctx.inner.QueryArgs().VisitAll(func(key, value []byte) {
		params[type_helper.BytesToString(key)] = NewParam(type_helper.BytesToString(value))
	})

	return params
}

func (ctx *Context) Header(name string) string {
	return type_helper.BytesToString(ctx.inner.Request.Header.Peek(name))
}

func (ctx *Context) Headers() map[string]string {
	requestHeaders := make(map[string]string)

	ctx.inner.Request.Header.VisitAll(func(key, value []byte) {
		requestHeaders[type_helper.BytesToString(key)] = type_helper.BytesToString(value)
	})

	return requestHeaders
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

	if reflect.ValueOf(object).Kind() != reflect.Ptr {
		return ErrorPointerRequired()
	}

	switch contentType {
	case content_types.JSON:
		if err := json.Unmarshal(ctx.Body(), &object); err != nil {
			return ErrorParseBody(err, content_types.JSON)
		}

		if err := ctx.validate.Struct(object); err != nil {
			return err
		}

		return nil
	case content_types.XML:
		if err := xml.Unmarshal(ctx.Body(), &object); err != nil {
			return ErrorParseBody(err, content_types.XML)
		}

		if err := ctx.validate.Struct(object); err != nil {
			return err
		}

		return nil
	}

	return ErrorUnknownContentType(contentType)
}

func (ctx *Context) Validate(object any) error {
	return ctx.validate.Struct(object)
}

func (ctx *Context) FormFile(key string) []byte {
	file, err := ctx.inner.FormFile(key)
	if err != nil {
		log.Error(err, "Parse form-data file")
		return nil
	}

	fileObject, err := file.Open()
	defer func() {
		_ = fileObject.Close()
	}()

	fileInBytes, err := io.ReadAll(fileObject)
	if err != nil {
		return nil
	}

	return fileInBytes
}

func (ctx *Context) IsWebSocket() bool {
	headerUpgrade := type_helper.BytesToString(ctx.inner.Request.Header.Peek("Upgrade"))
	return strings.EqualFold(headerUpgrade, "websocket")
}

func (ctx *Context) IsTLS() bool {
	return ctx.inner.IsTLS()
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

func (ctx *Context) SetHeader(key, value string) interfaces.Context {
	ctx.Response().Header.Set(key, value)
	return ctx
}

func (ctx *Context) SetCookie(key, value string) interfaces.Context {
	newCookie := &fasthttp.Cookie{}
	newCookie.SetKey(key)
	newCookie.SetValue(value)
	ctx.Response().Header.SetCookie(newCookie)
	return ctx
}

func (ctx *Context) Status(status int) interfaces.Context {
	ctx.status = status
	return ctx
}

func (ctx *Context) Empty() error {
	ctx.writer.Write(content_types.Text, ctx.status, nil)

	return nil
}

func (ctx *Context) String(message string) error {
	ctx.writer.Write(content_types.Text, ctx.status, type_helper.StringToBytes(message))

	return nil

}

func (ctx *Context) Bytes(body []byte) error {
	ctx.writer.Write(content_types.Bytes, ctx.status, body)

	return nil
}

func (ctx *Context) JSON(body any) error {
	bodyInBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ctx.writer.Write(content_types.JSON, ctx.status, bodyInBytes)

	return nil
}

func (ctx *Context) XML(body any) error {
	bodyInBytes, err := xml.Marshal(body)
	if err != nil {
		return err
	}

	ctx.writer.Write(content_types.XML, ctx.status, bodyInBytes)

	return nil
}

func (ctx *Context) Error(err error) error {
	if err == nil {
		return nil
	}

	boostError, ok := err.(interfaces.Error)
	if !ok {
		boostError = ErrorUnknownType(err)
	}

	ctx.writer.Write(
		boostError.ContentType(),
		boostError.HttpCode(),
		boostError.JSON(),
	)

	return err
}

func (ctx *Context) Redirect(url string, customStatus ...int) error {
	ctx.inner.Response.Header.SetCanonical(
		type_helper.StringToBytes("Location"),
		type_helper.StringToBytes(url),
	)

	if len(customStatus) > 0 {
		ctx.Status(customStatus[0])
	}

	ctx.inner.Redirect(url, ctx.status)

	return nil
}

func (ctx *Context) Next() error {
	nextHandler := ctx.nextHandler

	// check need to load action
	if !ctx.actionCalled.Load() && ctx.goingToCallAction.Load() {
		ctx.actionCalled.Store(true)
		err := ctx.action(ctx)
		if err != nil {
			return err
		}

		return nil
	}

	// if action already called
	if ctx.actionCalled.Load() {
		return nil
	}

	ctx.handlersChainIndex++

	if ctx.handlersChainIndex >= len(ctx.handlersChain) {
		ctx.nextHandler = ctx.action
		ctx.goingToCallAction.Store(true)
	} else {
		ctx.nextHandler = ctx.handlersChain[ctx.handlersChainIndex]
	}

	if nextHandler == nil {
		return nil
	}

	return nextHandler(ctx)
}

func (ctx *Context) Context() context.Context {
	if ctx.userCtx == nil {
		return context.Background()
	}

	return ctx.userCtx
}

func (ctx *Context) SetContext(userContext context.Context) {
	ctx.userCtx = userContext
}

func (ctx *Context) SetPanicHandler(panicHandler func(err error)) {
	ctx.panicHandler = panicHandler
}

func (ctx *Context) PanicHandler() func(err error) {
	return ctx.panicHandler
}
