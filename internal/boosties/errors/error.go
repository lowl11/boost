package errors

import (
	"encoding/json"
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
		grpcCode:  toGrpcCode(output.Code),
		context:   output.Context,
	}, true
}

func toGrpcCode(code int) codes.Code {
	switch code {
	case http.StatusOK:
		return codes.OK
	case http.StatusNotFound:
		return codes.NotFound
	case 499:
		return codes.Canceled
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return codes.InvalidArgument
	case http.StatusInternalServerError:
		return codes.Internal
	}

	return codes.Unknown
}

func toHttpCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Canceled:
		return 499
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Aborted:
		return http.StatusConflict
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Internal, codes.DataLoss, codes.Unknown:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
