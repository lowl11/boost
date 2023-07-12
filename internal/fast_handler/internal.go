package fast_handler

import (
	"fmt"
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/lowl11/boost/internal/helpers/fast_helper"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/content_types"
	"github.com/valyala/fasthttp"
	"net/http"
)

const (
	methodAny = "ANY"
)

func (handler *Handler) commonHandler(ctx *fasthttp.RequestCtx) {
	routeCtx, ok := handler.router.Search(type_helper.BytesToString(ctx.Path()))

	// if route not found
	if !ok {
		fast_helper.Write(
			ctx,
			content_types.JSON,
			http.StatusNotFound,
			type_helper.StringToBytes("{\"message\": \"not found\"}"),
		)
		// TODO: implement normal NOT FOUND object
		return
	}

	// if incoming method & registered are not match
	// in other case, registered may use method "ANY"
	if routeCtx.Method != type_helper.BytesToString(ctx.Method()) && routeCtx.Method != methodAny {
		fast_helper.Write(
			ctx,
			content_types.JSON,
			http.StatusMethodNotAllowed,
			type_helper.StringToBytes("{\"message\": \"method not allowed\"}"),
		)
		return
	}

	// call action
	err := routeCtx.Action(context.New(ctx).SetParams(routeCtx.Params))
	if err != nil {
		// TODO: implement me
		fmt.Println("handler action error:", err)
	}
}
