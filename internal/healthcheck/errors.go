package healthcheck

import (
	"github.com/lowl11/boost/errors"
	"net/http"
)

const (
	typeErrorHealthcheck = "HealthcheckError"
)

func errorHealthcheck(err error) error {
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	} else {
		errorMessage = "Healthcheck error"
	}

	return errors.
		New(errorMessage).
		SetType(typeErrorHealthcheck).
		SetHttpCode(http.StatusBadGateway)
}
