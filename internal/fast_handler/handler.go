package fast_handler

import (
	"github.com/lowl11/boost/internal/boosties/router"
	"github.com/lowl11/boost/pkg/types"
)

type Handler struct {
	router            *router.Router
	globalMiddlewares []types.HandlerFunc
	groupMiddlewares  map[string][]types.HandlerFunc
}

func New() *Handler {
	return &Handler{
		router:            router.New(),
		globalMiddlewares: make([]types.HandlerFunc, 0),
		groupMiddlewares:  make(map[string][]types.HandlerFunc),
	}
}
