package exception

func Try(tryFunc func() error) (err error) {
	defer func() {
		err = CatchPanic(recover())
	}()

	if err = tryFunc(); err != nil {
		return err
	}

	return nil
}
