package fast_handler

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

func (handler *Handler) commonHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/hello":
		_, _ = ctx.Write([]byte("hello world"))
	default:
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetBody([]byte("{\"message\": \"error message\"}"))
	}
}
