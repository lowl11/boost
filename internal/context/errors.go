package context

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"net/http"
)

const (
	typeErrorUnknownType        = "UnknownType"
	typeErrorUnknownContentType = "UnknownContentType"
	typeErrorParseBody          = "ParseBody"
	typeErrorPointerRequired    = "PointerRequired"
)

func ErrorParseIntParam(err error, value string) error {
	return errors.
		New("Parse int param error").
		SetType("ParseIntParamError").
		SetHttpCode(http.StatusUnprocessableEntity).
		SetError(err).
		AddContext("value", value)
}

func ErrorParseUUIDParam(err error, value string) error {
	return errors.
		New("Parse UUID param error").
		SetType("ParseUUIDParamError").
		SetHttpCode(http.StatusUnprocessableEntity).
		SetError(err).
		AddContext("value", value)
}

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

func ErrorPointerRequired() error {
	return errors.
		New("Pointer required").
		SetType(typeErrorPointerRequired).
		SetHttpCode(http.StatusInternalServerError)
}
