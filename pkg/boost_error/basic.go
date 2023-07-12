package boost_error

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/interfaces"
	"net/http"
)

const (
	TypeErrorUnknown          = "Unknown error"
	TypeErrorEndpointNotFound = "Endpoint not found"
	TypeErrorMethodNotAllowed = "Method not allowed"
)

func ErrorUnknown(err error) interfaces.Error {
	return errors.
		New("Unknown error: " + err.Error()).
		SetType(TypeErrorUnknown).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorNotFound() interfaces.Error {
	return errors.
		New("Endpoint not found").
		SetType(TypeErrorEndpointNotFound).
		SetHttpCode(http.StatusNotFound)
}

func ErrorMethodNotAllowed() interfaces.Error {
	return errors.
		New("Method not allowed").
		SetType(TypeErrorMethodNotAllowed).
		SetHttpCode(http.StatusMethodNotAllowed)
}
