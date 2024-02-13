package validator

import (
	"github.com/lowl11/boost/errors"
	"net/http"
)

func ErrorModelValidation(validate ...string) error {
	return errors.
		New("Model validation error").
		SetType(errorType).
		SetHttpCode(http.StatusUnprocessableEntity).
		AddContext("validate", validate)
}
