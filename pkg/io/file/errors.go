package file

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

func ErrorAlreadyDestroyed() interfaces.Error {
	return errors.
		New("File already destroyed").
		SetType("FileAlreadyDestroyed")
}

func ErrorNotFound(names ...string) interfaces.Error {
	return errors.
		New("File not exist").
		SetHttpCode(http.StatusNotFound).
		SetType("FileNotFound").
		AddContext("names", names)
}

func ErrorFileIsFolder() interfaces.Error {
	return errors.
		New("File is folder").
		SetType("FileIsFolder")
}
