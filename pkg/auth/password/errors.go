package password

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorPasswordsNotEqual = "PasswordsNotEqual"
)

func ErrorPasswordsNotEqual() error {
	return errors.
		New("Passwords are not equal").
		SetType(typeErrorPasswordsNotEqual).
		SetHttpCode(http.StatusUnprocessableEntity)
}
