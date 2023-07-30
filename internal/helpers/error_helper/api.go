package error_helper

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

func ToGrpcCode(code int) codes.Code {
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
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusInternalServerError:
		return codes.Internal
	}

	return codes.Unknown
}

func ToHttpCode(code codes.Code) int {
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
	case codes.InvalidArgument:
	case codes.FailedPrecondition:
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Internal:
	case codes.DataLoss:
	case codes.Unknown:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
