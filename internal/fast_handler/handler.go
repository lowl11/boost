package fast_handler

import (
	"github.com/lowl11/boost/internal/fast_handler/counter"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boost/pkg/system/validator"
	"github.com/valyala/fasthttp"
)

type CorsConfig struct {
	Enabled bool
	Origin  string
	Headers []string
	Methods []string
	Vary    []string

	debugPrint bool
}

type Handler struct {
	server            *fasthttp.Server
	router            *router
	globalMiddlewares []types.HandlerFunc
	groupMiddlewares  map[string][]types.HandlerFunc
	counter           *counter.Counter
	validate          *validator.Validator
	corsConfig        CorsConfig
	panicHandler      types.PanicHandler
}

func New(validate *validator.Validator) *Handler {
	return &Handler{
		server:            getServer(),
		router:            newRouter(),
		globalMiddlewares: make([]types.HandlerFunc, 0),
		groupMiddlewares:  make(map[string][]types.HandlerFunc),
		counter:           counter.New(),
		validate:          validate,
	}
}
