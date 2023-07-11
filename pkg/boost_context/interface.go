package boost_context

import "github.com/valyala/fasthttp"

type Context interface {
	Request() *fasthttp.Request

	Status(status int) Context

	JSON(body any) error
	String(message string) error
}
