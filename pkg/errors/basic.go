package errors

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/interfaces"
	"net/http"
)

const (
	TypeErrorUnknown          = "UnknownError"
	TypeErrorPanic            = "PanicError"
	TypeErrorRouteNotFound    = "RouteNotFound"
	TypeErrorMethodNotAllowed = "MethodNotAllowed"
)

// ErrorUnknown returns Boost Error for unknown type errors
func ErrorUnknown(err error) interfaces.Error {
	return errors.
		New("Unknown error: " + err.Error()).
		SetType(TypeErrorUnknown).
		SetHttpCode(http.StatusInternalServerError)
}

// ErrorPanic returns Boost Error for panics
func ErrorPanic(err error) interfaces.Error {
	return errors.
		New("PANIC RECOVER: " + err.Error()).
		SetType(TypeErrorPanic).
		SetHttpCode(http.StatusInternalServerError)
}

// ErrorNotFound returns Boost Error for not found endpoints
func ErrorNotFound() interfaces.Error {
	return errors.
		New("Route not found").
		SetType(TypeErrorRouteNotFound).
		SetHttpCode(http.StatusNotFound)
}

// ErrorMethodNotAllowed returns Boost Error for not allowed request method
func ErrorMethodNotAllowed() interfaces.Error {
	return errors.
		New("Method not allowed").
		SetType(TypeErrorMethodNotAllowed).
		SetHttpCode(http.StatusMethodNotAllowed)
}
