package context

import "github.com/valyala/fasthttp"

type IBoostContext interface {
	Request() *fasthttp.Request
	
	Status(status int) IBoostContext

	JSON(body any) error
	String(message string) error
}
