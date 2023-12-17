package try

import (
	"github.com/lowl11/boost/data/interfaces"
)

func (try *Try) Do() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		callCatch(try.catchFunc, err)
	}()

	try.tryFunc()
	if try.finallyFunc != nil {
		try.finallyFunc()
	}
}

func (try *Try) Catch(catch func(err error)) interfaces.Try {
	try.catchFunc = catch
	return try
}

func (try *Try) Finally(finally func()) interfaces.Try {
	try.finallyFunc = finally
	return try
}
