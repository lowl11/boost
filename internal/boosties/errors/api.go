package errors

func (err *Error) HttpCode(code int) *Error {
	err.httpCode = code
	return err
}

func (err *Error) Type(errorType string) *Error {
	err.errorType = errorType
	return err
}
