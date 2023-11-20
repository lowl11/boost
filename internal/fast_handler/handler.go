package fast_handler

import (
	"github.com/lowl11/boost/internal/boosties/router"
	"github.com/lowl11/boost/internal/services/boost/counter"
	"github.com/lowl11/boost/internal/services/system/validator"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/valyala/fasthttp"
)

type CorsConfig struct {
	Enabled bool
	Origin  string
	Headers []string
	Methods []string
	Vary    []string
}

type Handler struct {
	server            *fasthttp.Server
	router            *router.Router
	globalMiddlewares []types.HandlerFunc
	groupMiddlewares  map[string][]types.HandlerFunc
	counter           *counter.Counter
	validate          *validator.Validator
	corsConfig        CorsConfig
}

func New(validate *validator.Validator) *Handler {
	return &Handler{
		server:            getServer(),
		router:            router.New(),
		globalMiddlewares: make([]types.HandlerFunc, 0),
		groupMiddlewares:  make(map[string][]types.HandlerFunc),
		counter:           counter.New(),
		validate:          validate,
	}
}
