package password

import (
	"github.com/lowl11/boost/errors"
	"net/http"
)

func ErrorPasswordsNotEqual() error {
	return errors.New("Passwords are not equal").
		SetType("Password_PasswordsNotEqual").
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorEncryptPassword(err error) error {
	return errors.New("Encrypt password error").
		SetType("Password_EncryptError").
		SetError(err)
}

func ErrorDecryptPassword(err error) error {
	return errors.New("Decrypt password error").
		SetType("DecryptPassword").
		SetError(err)
}

func ErrorEncryptedPasswordEmpty() error {
	return errors.New("Encrypted password is empty").
		SetType("EncryptedPasswordEmpty")
}

func ErrorDecryptedPasswordEmpty() error {
	return errors.New("Decrypted password is empty").
		SetType("DecryptedPasswordEmpty")
}
