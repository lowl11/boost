package fast_handler

import (
	"github.com/lowl11/boost/internal/boosties/router"
	"github.com/lowl11/boost/internal/services/counter"
	"github.com/lowl11/boost/pkg/types"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	server            *fasthttp.Server
	router            *router.Router
	globalMiddlewares []types.HandlerFunc
	groupMiddlewares  map[string][]types.HandlerFunc
	counter           *counter.Counter
}

func New() *Handler {
	return &Handler{
		server:            getServer(),
		router:            router.New(),
		globalMiddlewares: make([]types.HandlerFunc, 0),
		groupMiddlewares:  make(map[string][]types.HandlerFunc),
		counter:           counter.New(),
	}
}
