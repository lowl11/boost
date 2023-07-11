package boost_context

import "github.com/valyala/fasthttp"

type Context interface {
	Request() *fasthttp.Request
	IsWebSocket() bool
	Param(name string) string
	QueryParam(name string) string

	Get(key string) any
	Set(key string, value any)

	Status(status int) Context

	Empty() error
	JSON(body any) error
	XML(body any) error
	String(message string) error
}
