package try

type Try struct {
	tryFunc     func()
	catchFunc   func(err error)
	finallyFunc func()
}

func New(tryFunc func()) *Try {
	return &Try{
		tryFunc: tryFunc,
	}
}
