package errors

import (
	"github.com/lowl11/boost/pkg/content_types"
	"strings"
)

func (err *Error) SetHttpCode(code int) *Error {
	err.httpCode = code
	return err
}

func (err *Error) HttpCode() int {
	return err.httpCode
}

func (err *Error) SetType(errorType string) *Error {
	err.errorType = errorType
	return err
}

func (err *Error) Type() string {
	return err.errorType
}

func (err *Error) ContentType() string {
	return content_types.JSON
}

func (err *Error) Error() string {
	builder := strings.Builder{}
	builder.Grow(len(err.message))

	builder.WriteString(err.message)

	return builder.String()
}
