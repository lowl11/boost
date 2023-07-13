package interfaces

type Route interface {
	Use(middlewareFunc ...func(context Context) error)
}
