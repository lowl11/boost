package errors

import (
	"encoding/json"
	"github.com/lowl11/boost/internal/helpers/error_helper"
	"google.golang.org/grpc/codes"
	"net/http"
)

const (
	unknownErrorType = "Unknown error type"
)

type Error struct {
	message    string
	errorType  string
	httpCode   int
	grpcCode   codes.Code
	context    map[string]any
	innerError error
}

type OutputError struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Type    string         `json:"type"`
	Code    int            `json:"code"`
	Context map[string]any `json:"context,omitempty"`
}

func New(message string) *Error {
	return &Error{
		message:   message,
		errorType: unknownErrorType,
		httpCode:  http.StatusInternalServerError,
		grpcCode:  codes.Unknown,
		context:   make(map[string]any),
	}
}

func Parse(response []byte) (*Error, bool) {
	output := OutputError{}

	if err := json.Unmarshal(response, &output); err != nil {
		return nil, false
	}

	return &Error{
		message:   output.Message,
		errorType: output.Type,
		httpCode:  output.Code,
		grpcCode:  error_helper.ToGrpcCode(output.Code),
		context:   output.Context,
	}, true
}
