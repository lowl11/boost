package interfaces

type Try interface {
	Catch(func(err error)) Try
	Finally(func()) Try
	Do()
}
