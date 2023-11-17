package fast_handler

import (
	"github.com/lowl11/boost/data/errors"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/context"
	"github.com/lowl11/boost/internal/boosties/panicer"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/config"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
)

const (
	methodAny      = "ANY"
	defaultMethods = "GET,POST,PUT,DELETE,OPTIONS,HEAD"
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

	if handler.corsConfig.Enabled {
		// fill CORS headers
		ctx.Response.Header.Set("Access-Control-Allow-Origin", handler.getOrigin(ctx))
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", handler.getHeaders(ctx))
		ctx.Response.Header.Set("Access-Control-Allow-Methods", handler.getMethods())

		// OPTIONS case
		if ctx.IsOptions() {
			ctx.SetStatusCode(http.StatusNoContent)
			return
		}
	}

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

func (handler *Handler) getOrigin(ctx *fasthttp.RequestCtx) string {
	origins := make([]string, 0, 3)
	requestOrigin := types.ToString(ctx.Request.Header.Peek("Origin"))
	if requestOrigin != "" {
		origins = append(origins, requestOrigin)
	}

	if handler.corsConfig.Origin != "" {
		origins = append(origins, handler.corsConfig.Origin)
	}

	if len(origins) > 0 {
		return strings.Join(origins, ",")
	}

	return "*"
}

func (handler *Handler) getHeaders(ctx *fasthttp.RequestCtx) string {
	accessHeaders := make([]string, 0, 10)
	for _, header := range ctx.Request.Header.PeekKeys() {
		accessHeaders = append(accessHeaders, types.ToString(header))
	}

	accessHeaders = append(accessHeaders, "Content-Type", "Authorization", "Origin")
	if len(handler.corsConfig.Methods) > 0 {
		for _, header := range handler.corsConfig.Headers {
			if header == "" {
				continue
			}

			accessHeaders = append(accessHeaders, header)
		}
	}
	return strings.Join(accessHeaders, ",")
}

func (handler *Handler) getMethods() string {
	// custom methods
	if len(handler.corsConfig.Methods) > 0 {
		var methods string
		for index, method := range handler.corsConfig.Methods {
			if method == "" {
				continue
			}

			methods += method
			if index < len(handler.corsConfig.Methods)-1 {
				methods += ","
			}
		}

		if len(methods) == 0 {
			return defaultMethods
		}

		return methods
	}

	return defaultMethods
}

func (handler *Handler) tryUpdateCORS() {
	if !handler.corsConfig.Enabled {
		handler.corsConfig.Enabled = strings.ToLower(config.Get("CORS_ENABLED")) == "true"
	}

	if handler.corsConfig.Origin == "" {
		handler.corsConfig.Origin = config.Get("CORS_ORIGIN")
	}

	if len(handler.corsConfig.Headers) == 0 {
		handler.corsConfig.Headers = strings.Split(config.Get("CORS_HEADERS"), ",")
	}

	if len(handler.corsConfig.Methods) == 0 {
		handler.corsConfig.Methods = strings.Split(config.Get("CORS_METHODS"), ",")
	}
}
