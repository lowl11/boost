package try

import "errors"

func callCatch(catchFunc func(err error), err any) {
	errStr, ok := err.(string)
	if ok {
		catchFunc(errors.New(errStr))
		return
	}

	errError, ok := err.(error)
	if ok {
		catchFunc(errError)
		return
	}
}
