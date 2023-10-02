package fast_handler

import (
	"github.com/lowl11/boost/data/errors"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/lowl11/boost/internal/boosties/panicer"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/lazylog/log"
	"github.com/valyala/fasthttp"
)

const (
	methodAny = "ANY"
)

func getServer() *fasthttp.Server {
	return &fasthttp.Server{
		ErrorHandler: writeUnknownError,
	}
}

func (handler *Handler) handler(ctx *fasthttp.RequestCtx) {
	// handler panic
	defer func() {
		err := panicer.Handle(recover())
		if err == nil {
			return
		}

		log.Error(err, "PANIC RECOVERED")
		writePanicError(ctx, err)
	}()

	// find route
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

	// get group middlewares
	groupMiddlewares, ok := handler.groupMiddlewares[routeCtx.GroupID]
	if !ok || routeCtx.GroupID == "" {
		groupMiddlewares = []types.HandlerFunc{}
	}

	// get endpoint middlewares
	endpointMiddlewares := routeCtx.Middlewares
	if endpointMiddlewares == nil {
		endpointMiddlewares = []types.HandlerFunc{}
	}

	// create handlers chain (with middlewares)
	// order which given handlers is IMPORTANT!!!
	handlersChain := append(handler.globalMiddlewares, groupMiddlewares...)
	handlersChain = append(handlersChain, endpointMiddlewares...)

	// create new boost context
	boostCtx := context.
		New(ctx, routeCtx.Action, handlersChain, handler.validate).
		SetParams(routeCtx.Params)

	// call chain of handlers/middlewares
	err := boostCtx.Next()
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

func writePanicError(ctx *fasthttp.RequestCtx, err error) {
	writeError(ctx, errors.ErrorPanic(err))
}

func writeError(ctx *fasthttp.RequestCtx, err interfaces.Error) {
	ctx.SetStatusCode(err.HttpCode())
	ctx.Response.Header.Set("Content-Type", err.ContentType())
	ctx.SetBody(type_helper.StringToBytes(err.Error()))
}
