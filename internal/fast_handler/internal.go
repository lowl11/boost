package fast_handler

import (
	"fmt"
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/valyala/fasthttp"
	"net/http"
)

const (
	methodAny = "ANY"
)

func (handler *Handler) commonHandler(ctx *fasthttp.RequestCtx) {
	routeCtx, ok := handler.router.Search(string(ctx.Path()))

	// if route not found
	if !ok {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetBody([]byte("{\"message\": \"not found\"}"))
		// TODO: implement normal NOT FOUND object
		return
	}

	// if incoming method & registered are not match
	// in other case, registered may use method "ANY"
	if routeCtx.Method != string(ctx.Method()) && routeCtx.Method != methodAny {
		ctx.SetStatusCode(http.StatusMethodNotAllowed)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetBody([]byte("{\"message\": \"method not allowed\"}"))
		return
	}

	// call action
	err := routeCtx.Action(context.New(ctx).SetParams(routeCtx.Params))
	if err != nil {
		// TODO: implement me
		fmt.Println("handler action error:", err)
	}
}
