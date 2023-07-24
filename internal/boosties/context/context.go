package context

import (
	"github.com/lowl11/boost/pkg/types"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"sync/atomic"
)

type Context struct {
	inner *fasthttp.RequestCtx

	status       int
	keyContainer sync.Map
	params       map[string]string

	action            types.HandlerFunc
	goingToCallAction atomic.Bool
	actionCalled      atomic.Bool

	nextHandler        types.HandlerFunc
	handlersChain      []types.HandlerFunc
	handlersChainIndex int
}

func New(inner *fasthttp.RequestCtx, action types.HandlerFunc, handlersChain []types.HandlerFunc) *Context {
	var nextHandler types.HandlerFunc
	if len(handlersChain) > 0 {
		nextHandler = handlersChain[0]
	}

	return &Context{
		inner:  inner,
		status: http.StatusOK,
		params: make(map[string]string),

		action:        action,
		nextHandler:   nextHandler,
		handlersChain: handlersChain,
	}
}
