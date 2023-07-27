package context

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/interfaces"
	"net/http"
)

const (
	typeErrorUnknownType        = "UnknownType"
	typeErrorUnknownContentType = "UnknownContentType"
	typeErrorParseBody          = "ParseBody"
)

func ErrorUnknownType(err error) interfaces.Error {
	return errors.
		New("Unknown error: " + err.Error()).
		SetType(typeErrorUnknownType).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorUnknownContentType(contentType string) interfaces.Error {
	return errors.
		New("Unknown Content-Type").
		SetType(typeErrorUnknownContentType).
		SetHttpCode(http.StatusBadRequest).
		AddContext("Content-Type", contentType)
}

func ErrorParseBody(err error, format string) error {
	return errors.
		New("Parse body error").
		SetType(typeErrorParseBody).
		SetHttpCode(http.StatusBadRequest).
		SetError(err).
		AddContext("format", format)
}
