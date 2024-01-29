package errors

import (
	"encoding/json"
	"github.com/lowl11/boost/data/enums/content_types"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/types"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
)

const (
	unknownErrorType = "Unknown error type"
)

type customError struct {
	message    string
	errorType  string
	httpCode   int
	grpcCode   codes.Code
	context    map[string]any
	innerError error
}

type outputError struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Type    string         `json:"type"`
	Code    int            `json:"code"`
	Context map[string]any `json:"context,omitempty"`
}

func newError(message string) *customError {
	return &customError{
		message:   message,
		errorType: unknownErrorType,
		httpCode:  http.StatusInternalServerError,
		grpcCode:  codes.Unknown,
		context:   make(map[string]any),
	}
}

func parseCustom(response []byte) (*customError, bool) {
	output := outputError{}

	if err := json.Unmarshal(response, &output); err != nil {
		return nil, false
	}

	return &customError{
		message:   output.Message,
		errorType: output.Type,
		httpCode:  output.Code,
		grpcCode:  toGrpcCode(output.Code),
		context:   output.Context,
	}, true
}

const (
	status = "ERROR"
)

func (err *customError) Message() string {
	return err.message
}

func (err *customError) SetHttpCode(code int) interfaces.Error {
	err.httpCode = code
	err.grpcCode = toGrpcCode(code)

	return err
}

func (err *customError) HttpCode() int {
	return err.httpCode
}

func (err *customError) GrpcCode() codes.Code {
	return err.grpcCode
}

func (err *customError) SetType(errorType string) interfaces.Error {
	err.errorType = errorType
	return err
}

func (err *customError) Type() string {
	return err.errorType
}

func (err *customError) ContentType() string {
	return content_types.JSON
}

func (err *customError) Context() map[string]any {
	return err.context
}

func (err *customError) SetContext(context map[string]any) interfaces.Error {
	for key, value := range context {
		err.context[key] = value
	}

	return err
}

func (err *customError) AddContext(key string, value any) interfaces.Error {
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

func (err *customError) InnerError() error {
	return err.innerError
}

func (err *customError) SetError(innerError error) interfaces.Error {
	err.innerError = innerError
	return err
}

func (err *customError) Error() string {
	errorMessage := err.message
	if err.innerError != nil {
		errorMessage += " | " + err.innerError.Error()
	}

	output := outputError{
		Status:  status,
		Message: errorMessage,
		Type:    err.errorType,
		Code:    err.httpCode,
		Context: err.context,
	}

	outputInBytes, _ := json.Marshal(output)
	return types.BytesToString(outputInBytes)
}

func (err *customError) JSON() []byte {
	errorMessage := err.message
	if err.innerError != nil {
		errorMessage += " | " + err.innerError.Error()
	}

	output := outputError{
		Status:  status,
		Message: errorMessage,
		Type:    err.errorType,
		Code:    err.httpCode,
		Context: err.context,
	}

	outputInBytes, _ := json.Marshal(output)
	return outputInBytes
}

func (err *customError) String() string {
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

func (err *customError) Is(compare error) bool {
	boostError, ok := compare.(interfaces.Error)
	if !ok {
		return false
	}

	return equals(err, boostError)
}

func equals(left, right interfaces.Error) bool {
	if left.HttpCode() == right.HttpCode() && left.Type() == right.Type() &&
		left.Error() == right.Error() {
		return true
	}

	return false
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
