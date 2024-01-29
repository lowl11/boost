package msgbus

import (
	"github.com/lowl11/boost/errors"
	"net/http"
)

const (
	typeErrorGetNameOfEvent = "GetNameOfEvent"
)

func ErrorGetNameOfEvent(err error) error {
	return errors.
		New("Get name of event error").
		SetType(typeErrorGetNameOfEvent).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
