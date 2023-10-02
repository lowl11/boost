package folder

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

func ErrorNotFound() interfaces.Error {
	return errors.
		New("Folder not found").
		SetType("FolderNotFound").
		SetHttpCode(http.StatusNotFound)
}
