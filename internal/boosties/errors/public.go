package errors

import (
	"encoding/json"
	"github.com/lowl11/boost/data/enums/content_types"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/types"
	"google.golang.org/grpc/codes"
	"strings"
)

const (
	status = "ERROR"
)

func (err *Error) Message() string {
	return err.message
}

func (err *Error) SetHttpCode(code int) interfaces.Error {
	err.httpCode = code
	err.grpcCode = toGrpcCode(code)

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
	if value == nil {
		return err
	}

	if arr, ok := value.([]string); ok {
		if len(arr) == 0 {
			return err
		}
	}

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
	return types.BytesToString(outputInBytes)
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

	if err.innerError != nil {
		builder.WriteString(". ")
		builder.WriteString(err.innerError.Error())
	}

	if err.context != nil && len(err.context) > 0 {
		builder.WriteString(". Context: ")
		for key, value := range err.context {
			if key == "trace" {
				trace := value.([]string)

				for _, traceLine := range trace {
					builder.WriteString("\n\t")
					builder.WriteString(traceLine)
				}
				continue
			}

			builder.WriteString(key)
			builder.WriteString("=")
			builder.WriteString(types.ToString(value))
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
