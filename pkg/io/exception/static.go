package exception

import "github.com/lowl11/boost/internal/boosties/panicer"

func Try(tryFunc func() error) (err error) {
	defer func() {
		err = panicer.Handle(recover())
	}()

	if err = tryFunc(); err != nil {
		return err
	}

	return nil
}

func CatchPanic(err any) error {
	return panicer.Handle(err)
}

func StackTrance() []string {
	return panicer.StackTrace()
}
