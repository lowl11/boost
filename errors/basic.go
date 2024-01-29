package errors

import (
	"github.com/lowl11/boost/data/interfaces"
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
	return New("Unknown error: " + err.Error()).
		SetType(TypeErrorUnknown).
		SetHttpCode(http.StatusInternalServerError)
}

// ErrorPanic returns Boost Error for panics
func ErrorPanic(err error) interfaces.Error {
	if boostError, ok := err.(interfaces.Error); ok {
		return boostError
	}

	return New("PANIC RECOVER: " + err.Error()).
		SetType(TypeErrorPanic).
		SetHttpCode(http.StatusInternalServerError)
}

// ErrorNotFound returns Boost Error for not found endpoints
func ErrorNotFound() interfaces.Error {
	return New("Route not found").
		SetType(TypeErrorRouteNotFound).
		SetHttpCode(http.StatusNotFound)
}

// ErrorMethodNotAllowed returns Boost Error for not allowed request method
func ErrorMethodNotAllowed() interfaces.Error {
	return New("Method not allowed").
		SetType(TypeErrorMethodNotAllowed).
		SetHttpCode(http.StatusMethodNotAllowed)
}
