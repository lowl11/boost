package socket

type HandlerFunc func(ctx *Context) error

type Handler struct {
	handler HandlerFunc
}

func NewHandler(handler HandlerFunc) *Handler {
	return &Handler{
		handler: handler,
	}
}

func (handler Handler) Run(conn *Conn, messageType int, body []byte) error {
	return handler.handler(newContext(conn, messageType, body))
}
