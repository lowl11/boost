package password

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorPasswordsNotEqual = "PasswordsNotEqual"
	typeErrorEncryptPassword   = "EncryptPassword"
	typeErrorDecryptPassword   = "DecryptPassword"
)

func ErrorPasswordsNotEqual() error {
	return errors.
		New("Passwords are not equal").
		SetType(typeErrorPasswordsNotEqual).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorEncryptPassword(err error) error {
	return errors.
		New("Encrypt password error").
		SetType(typeErrorEncryptPassword).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorDecryptPassword(err error) error {
	return errors.
		New("Decrypt password error").
		SetType(typeErrorDecryptPassword).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
