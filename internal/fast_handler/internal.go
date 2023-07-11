package fast_handler

import (
	"fmt"
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (handler *Handler) commonHandler(ctx *fasthttp.RequestCtx) {
	routeCtx, ok := handler.router.Map()[string(ctx.Path())]
	if !ok {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetBody([]byte("{\"message\": \"error message\"}"))
		// TODO: implement normal NOT FOUND object
		return
	}

	err := routeCtx.Action(context.New(ctx))
	if err != nil {
		// TODO: implement me
		fmt.Println("handler action error:", err)
	}
}
