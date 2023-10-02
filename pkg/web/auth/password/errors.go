package password

import (
	"github.com/lowl11/boost/data/errors"
	"net/http"
)

const (
	typeErrorPasswordsNotEqual      = "PasswordsNotEqual"
	typeErrorEncryptPassword        = "EncryptPassword"
	typeErrorDecryptPassword        = "DecryptPassword"
	typeErrorDecryptedPasswordEmpty = "DecryptedPasswordEmpty"
	typeErrorEncryptedPasswordEmpty = "EncryptedPasswordEmpty"
)

func ErrorPasswordsNotEqual() error {
	return errors.New("Passwords are not equal").
		SetType(typeErrorPasswordsNotEqual).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorEncryptPassword(err error) error {
	return errors.New("Encrypt password error").
		SetType(typeErrorEncryptPassword).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorDecryptPassword(err error) error {
	return errors.New("Decrypt password error").
		SetType(typeErrorDecryptPassword).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorEncryptedPasswordEmpty() error {
	return errors.New("Encrypted password is empty").
		SetType(typeErrorEncryptedPasswordEmpty).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorDecryptedPasswordEmpty() error {
	return errors.New("Decrypted password is empty").
		SetType(typeErrorDecryptedPasswordEmpty).
		SetHttpCode(http.StatusInternalServerError)
}
