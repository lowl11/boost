package exception

import "github.com/lowl11/boost/errors"

func CatchPanic(err any) error {
	if err == nil {
		return nil
	}

	parsedError := fromAny(err)
	return errors.New("PANIC RECOVER: "+parsedError).
		SetType("PanicError").
		AddContext("trace", StackTrace())
}

func fromAny(err any) string {
	if err == nil {
		return ""
	}

	switch err.(type) {
	case string:
		return err.(string)
	case error:
		return err.(error).Error()
	}

	return ""
}
