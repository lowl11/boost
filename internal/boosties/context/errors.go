package context

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/boost_error"
	"net/http"
)

const (
	typeErrorUnknownContentType = "UnknownContentType"
)

func ErrorUnknownContentType(contentType string) boost_error.Error {
	return errors.
		New("unknown content-type: " + contentType).
		SetType(typeErrorUnknownContentType).
		SetHttpCode(http.StatusBadRequest)
}
