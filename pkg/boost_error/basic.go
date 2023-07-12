package boost_error

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

const (
	TypeErrorUnknown          = "Unknown error"
	TypeErrorEndpointNotFound = "Endpoint not found"
	TypeErrorMethodNotAllowed = "Method not allowed"
)

func ErrorUnknown() Error {
	return errors.
		New("Unknown error").
		SetType(TypeErrorUnknown).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorNotFound() Error {
	return errors.
		New("Endpoint not found").
		SetType(TypeErrorEndpointNotFound).
		SetHttpCode(http.StatusNotFound)
}

func ErrorMethodNotAllowed() Error {
	return errors.
		New("Method not allowed").
		SetType(TypeErrorMethodNotAllowed).
		SetHttpCode(http.StatusMethodNotAllowed)
}
