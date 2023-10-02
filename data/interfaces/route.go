package interfaces

// Route is interface which will return after adding new route
type Route interface {
	Use(middlewareFunc ...func(context Context) error)
}
