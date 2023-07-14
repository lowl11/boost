package context

import (
	"github.com/lowl11/boost/pkg/types"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

type Context struct {
	inner *fasthttp.RequestCtx

	status       int
	keyContainer sync.Map
	params       map[string]string

	action             types.HandlerFunc
	nextHandler        types.HandlerFunc
	handlersChain      []types.HandlerFunc
	handlersChainIndex int
}

func New(inner *fasthttp.RequestCtx, action types.HandlerFunc, handlersChain []types.HandlerFunc) *Context {
	return &Context{
		inner:  inner,
		status: http.StatusOK,
		params: make(map[string]string),

		action:        action,
		nextHandler:   handlersChain[0],
		handlersChain: handlersChain,
	}
}
