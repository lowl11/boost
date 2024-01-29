package fast_handler

import (
	"fmt"
	"github.com/lowl11/boost/config"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/internal/context"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
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
	server := &fasthttp.Server{
		ErrorHandler: writeUnknownError,
	}
	server.MaxConnsPerIP = 10
	server.MaxRequestsPerConn = 10
	return &fasthttp.Server{
		ErrorHandler: writeUnknownError,
	}
}

func (handler *Handler) handler(ctx *fasthttp.RequestCtx) {
	var boostCtx *context.Context

	// handler panic
	defer func() {
		err := exception.CatchPanic(recover())
		if err == nil {
			return
		}

		log.Error(err, "PANIC RECOVERED")

		if handler.panicHandler != nil {
			handler.panicHandler(err)
		}

		if boostCtx.PanicHandler() != nil {
			boostCtx.PanicHandler()(err)
		}

		writePanicError(ctx, err)
	}()

	if handler.corsConfig.Enabled {
		// fill CORS headers
		ctx.Response.Header.Set("Access-Control-Allow-Origin", handler.getOrigin(ctx))
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", handler.getHeaders(ctx))
		ctx.Response.Header.Set("Access-Control-Allow-Methods", handler.getMethods())
		ctx.Response.Header.Set("Access-Control-Max-Age", "1728000")
		ctx.Response.Header.Set("Vary", handler.getVary())

		// OPTIONS case
		if ctx.IsOptions() {
			ctx.SetStatusCode(http.StatusNoContent)
			return
		}
	}

	// find route
	routeCtx, ok := handler.router.Search(types.BytesToString(ctx.Path()))

	// if route not found
	if !ok {
		writeError(ctx, errors.ErrorNotFound())
		return
	}

	// if incoming method & registered are not match
	// in other case, registered may use method "ANY"
	if routeCtx.Method != types.BytesToString(ctx.Method()) && routeCtx.Method != methodAny {
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
	boostCtx = context.New(ctx, routeCtx.Action, handlersChain, handler.validate).
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
	ctx.SetBody(types.StringToBytes(err.Error()))
}

func (handler *Handler) getOrigin(ctx *fasthttp.RequestCtx) string {
	if handler.corsConfig.debugPrint {
		headers := make(map[string]string)
		ctx.Request.Header.VisitAll(func(key, value []byte) {
			headers[types.ToString(key)] = types.ToString(value)
		})
		log.Info("CORS Debug -> Headers:", headers)
	}

	// get from config
	if handler.corsConfig.Origin != "" {
		if handler.corsConfig.debugPrint {
			log.Info("CORS origin from config:", handler.corsConfig.Origin)
		}
		return handler.corsConfig.Origin
	}

	// get from request headers
	requestOrigin := types.ToString(ctx.Request.Header.Peek("Origin"))
	if requestOrigin != "" {
		if handler.corsConfig.debugPrint {
			log.Info("CORS origin from request 'Origin' header:", requestOrigin)
		}
		return requestOrigin
	}

	// try to build dynamic
	dynamicOrigin := strings.Builder{}
	dynamicOrigin.Grow(len(ctx.URI().Scheme()) + len(ctx.URI().Host()) + 3)
	_, _ = fmt.Fprintf(&dynamicOrigin, "%s://%s", ctx.URI().Scheme(), ctx.URI().Host())
	if handler.corsConfig.debugPrint {
		if dynamicOrigin.String() != "" {
			log.Info("CORS origin build by dynamic URL context:", dynamicOrigin.String())
		}
	}
	return dynamicOrigin.String()
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

			var needContinue bool
			for _, accessHeader := range accessHeaders {
				if header == accessHeader {
					needContinue = true
					break
				}
			}

			if needContinue {
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
		handler.corsConfig.Enabled = config.Get("CORS_ENABLED").Bool()
	}

	if handler.corsConfig.Origin == "" {
		handler.corsConfig.Origin = config.Get("CORS_ORIGIN").String()
	}

	if len(handler.corsConfig.Headers) == 0 {
		handler.corsConfig.Headers = config.Get("CORS_HEADERS").Strings()
	}

	if len(handler.corsConfig.Methods) == 0 {
		handler.corsConfig.Methods = config.Get("CORS_METHODS").Strings()
	}

	if len(handler.corsConfig.Vary) == 0 {
		handler.corsConfig.Vary = config.Get("CORS_VARY").Strings()
	}

	handler.corsConfig.debugPrint = config.Get("CORS_DEBUG").Bool()
}

func (handler *Handler) getVary() string {
	if handler.corsConfig.Vary == nil || len(handler.corsConfig.Vary) == 0 {
		return "*"
	}

	var varyHeader string
	for index, header := range handler.corsConfig.Vary {
		varyHeader += header
		if index <= len(handler.corsConfig.Vary)-1 {
			varyHeader += ","
		}
	}

	return varyHeader
}
