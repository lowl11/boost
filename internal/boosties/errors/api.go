package errors

import (
	"encoding/json"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/pkg/content_types"
	"github.com/lowl11/boost/pkg/interfaces"
)

const (
	status = "ERROR"
)

func (err *Error) SetHttpCode(code int) interfaces.Error {
	err.httpCode = code
	return err
}

func (err *Error) HttpCode() int {
	return err.httpCode
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

func (err *Error) Error() string {
	output := OutputError{
		Status:  status,
		Message: err.message,
		Type:    err.errorType,
		Code:    err.httpCode,
	}

	outputInBytes, _ := json.Marshal(output)
	return type_helper.BytesToString(outputInBytes)
}

func (err *Error) JSON() []byte {
	output := OutputError{
		Status:  status,
		Message: err.message,
		Type:    err.errorType,
		Code:    err.httpCode,
	}

	outputInBytes, _ := json.Marshal(output)
	return outputInBytes
}
