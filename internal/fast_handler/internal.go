package fast_handler

import (
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/lowl11/boost/internal/helpers/fast_helper"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/boost_error"
	"github.com/valyala/fasthttp"
)

const (
	methodAny = "ANY"
)

func (handler *Handler) commonHandler(ctx *fasthttp.RequestCtx) {
	routeCtx, ok := handler.router.Search(type_helper.BytesToString(ctx.Path()))

	// if route not found
	if !ok {
		writeError(ctx, boost_error.ErrorNotFound())
		return
	}

	// if incoming method & registered are not match
	// in other case, registered may use method "ANY"
	if routeCtx.Method != type_helper.BytesToString(ctx.Method()) && routeCtx.Method != methodAny {
		writeError(ctx, boost_error.ErrorMethodNotAllowed())
		return
	}

	// call action
	err := routeCtx.Action(context.New(ctx).SetParams(routeCtx.Params))
	if err != nil {
		boostError, errorParse := err.(boost_error.Error)
		if !errorParse {
			writeUnknownError(ctx)
			return
		}

		writeError(ctx, boostError)

		return
	}
}

func writeUnknownError(ctx *fasthttp.RequestCtx) {
	writeError(ctx, boost_error.ErrorUnknown())
}

func writeError(ctx *fasthttp.RequestCtx, err boost_error.Error) {
	fast_helper.Write(
		ctx,
		err.ContentType(),
		err.HttpCode(),
		type_helper.StringToBytes(err.Error()),
	)
}
