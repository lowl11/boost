package fast_handler

import (
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/lowl11/boost/internal/helpers/fast_helper"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/errors"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/valyala/fasthttp"
)

const (
	methodAny = "ANY"
)

func (handler *Handler) commonHandler(ctx *fasthttp.RequestCtx) {
	routeCtx, ok := handler.router.Search(type_helper.BytesToString(ctx.Path()))

	// if route not found
	if !ok {
		writeError(ctx, errors.ErrorNotFound())
		return
	}

	// if incoming method & registered are not match
	// in other case, registered may use method "ANY"
	if routeCtx.Method != type_helper.BytesToString(ctx.Method()) && routeCtx.Method != methodAny {
		writeError(ctx, errors.ErrorMethodNotAllowed())
		return
	}

	// call action
	err := routeCtx.Action(context.New(ctx).SetParams(routeCtx.Params))
	if err != nil {
		boostError, errorParse := err.(interfaces.Error)
		if !errorParse {
			writeUnknownError(ctx, err)
			return
		}

		writeError(ctx, boostError)

		return
	}
}

func writeUnknownError(ctx *fasthttp.RequestCtx, err error) {
	writeError(ctx, errors.ErrorUnknown(err))
}

func writeError(ctx *fasthttp.RequestCtx, err interfaces.Error) {
	fast_helper.Write(
		ctx,
		err.ContentType(),
		err.HttpCode(),
		type_helper.StringToBytes(err.Error()),
	)
}
