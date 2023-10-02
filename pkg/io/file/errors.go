package file

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

func ErrorAlreadyDestroyed() interfaces.Error {
	return errors.
		New("File already destroyed")
}

func ErrorNotFound() interfaces.Error {
	return errors.
		New("File not exist").
		SetHttpCode(http.StatusNotFound)
}
