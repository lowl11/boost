package context

import (
	"context"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boost/pkg/system/validator"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"sync/atomic"
)

type Context struct {
	inner    *fasthttp.RequestCtx
	writer   *fastWriter
	validate *validator.Validator

	status       int
	keyContainer sync.Map
	params       map[string]string

	action            types.HandlerFunc
	goingToCallAction atomic.Bool
	actionCalled      atomic.Bool

	nextHandler        types.HandlerFunc
	handlersChain      []types.HandlerFunc
	handlersChainIndex int

	userCtx      context.Context
	panicHandler types.PanicHandler
}

func New(
	inner *fasthttp.RequestCtx,
	action types.HandlerFunc,
	handlersChain []types.HandlerFunc,
	validate *validator.Validator,
) *Context {
	var nextHandler types.HandlerFunc
	if len(handlersChain) > 0 {
		nextHandler = handlersChain[0]
	}

	return &Context{
		inner:    inner,
		writer:   newFastWriter(inner),
		validate: validate,

		status: http.StatusOK,
		params: make(map[string]string),

		action:        action,
		nextHandler:   nextHandler,
		handlersChain: handlersChain,
	}
}
