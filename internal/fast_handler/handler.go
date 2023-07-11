package fast_handler

import "github.com/lowl11/boost/internal/boosties/router"

type Handler struct {
	router *router.Router
}

func New() *Handler {
	return &Handler{
		router: router.New(),
	}
}
