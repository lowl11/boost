package errors

import (
	"encoding/json"
	"github.com/lowl11/boost/internal/helpers/error_helper"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/enums/content_types"
	"github.com/lowl11/boost/pkg/interfaces"
	"google.golang.org/grpc/codes"
	"strings"
)

const (
	status = "ERROR"
)

func (err *Error) SetHttpCode(code int) interfaces.Error {
	err.httpCode = code
	err.grpcCode = error_helper.ToGrpcCode(code)

	return err
}

func (err *Error) HttpCode() int {
	return err.httpCode
}

func (err *Error) GrpcCode() codes.Code {
	return err.grpcCode
}

func (err *Error) SetType(errorType string) interfaces.Error {
	err.errorType = errorType
	return err
}

func (err *Error) Type() string {
	return err.errorType
}

func (err *Error) ContentType() string {
	return content_types.JSON
}

func (err *Error) Context() map[string]any {
	return err.context
}

func (err *Error) SetContext(context map[string]any) interfaces.Error {
	for key, value := range context {
		err.context[key] = value
	}

	return err
}

func (err *Error) AddContext(key string, value any) interfaces.Error {
	err.context[key] = value

	return err
}

func (err *Error) InnerError() error {
	return err.innerError
}

func (err *Error) SetError(innerError error) interfaces.Error {
	err.innerError = innerError
	return err
}

func (err *Error) Error() string {
	errorMessage := err.message
	if err.innerError != nil {
		errorMessage += " | " + err.innerError.Error()
	}

	output := OutputError{
		Status:  status,
		Message: errorMessage,
		Type:    err.errorType,
		Code:    err.httpCode,
		Context: err.context,
	}

	outputInBytes, _ := json.Marshal(output)
	return type_helper.BytesToString(outputInBytes)
}

func (err *Error) JSON() []byte {
	errorMessage := err.message
	if err.innerError != nil {
		errorMessage += " | " + err.innerError.Error()
	}

	output := OutputError{
		Status:  status,
		Message: errorMessage,
		Type:    err.errorType,
		Code:    err.httpCode,
		Context: err.context,
	}

	outputInBytes, _ := json.Marshal(output)
	return outputInBytes
}

func (err *Error) String() string {
	builder := strings.Builder{}
	builder.Grow(500)
	builder.WriteString(err.message)
	builder.WriteString(". ")
	builder.WriteString("")

	if err.context != nil {
		builder.WriteString("Context: ")
		for key, value := range err.context {
			builder.WriteString(key)
			builder.WriteString("=")
			builder.WriteString(type_helper.ToString(value, false))
			builder.WriteString(";")
		}
	}

	return builder.String()
}

func (err *Error) Is(compare error) bool {
	boostError, ok := compare.(interfaces.Error)
	if !ok {
		return false
	}

	return equals(err, boostError)
}
