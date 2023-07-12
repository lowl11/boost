package context

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/interfaces"
	"net/http"
)

const (
	typeErrorUnknownType        = "UnknownType"
	typeErrorUnknownContentType = "UnknownContentType"
)

func ErrorUnknownType(err error) interfaces.Error {
	return errors.
		New("Unknown error: " + err.Error()).
		SetType(typeErrorUnknownType).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorUnknownContentType(contentType string) interfaces.Error {
	return errors.
		New("Unknown Content-Type: " + contentType).
		SetType(typeErrorUnknownContentType).
		SetHttpCode(http.StatusBadRequest)
}
