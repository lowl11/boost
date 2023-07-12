package boost_context

import "github.com/valyala/fasthttp"

type Context interface {
	Request() *fasthttp.Request
	Param(name string) string
	QueryParam(name string) string
	Header(name string) string
	Headers() map[string]string
	Cookie(name string) string
	Cookies() map[string]string
	Body() []byte
	Parse(object any) error

	IsWebSocket() bool

	Get(key string) any
	Set(key string, value any)

	Status(status int) Context

	Empty() error
	String(message string) error
	Bytes(body []byte) error
	JSON(body any) error
	XML(body any) error
}
