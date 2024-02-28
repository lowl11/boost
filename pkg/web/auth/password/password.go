package password

import (
	"github.com/lowl11/boost/pkg/system/types"
	"math/rand"
)

// Compare two passwords.
// If they are not equal returns boost error
func Compare(password, rePassword string) error {
	if password == rePassword {
		return nil
	}

	return ErrorPasswordsNotEqual()
}

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers        = "0123456789"
	specialSymbols = "!@#$%^&*()_+=-"
)

// Generate generates new password which can include special symbols and numbers
func Generate(length int, useLetters, includeSpecial, includeNumber bool) string {
	var password []byte
	var charSource string

	if includeNumber {
		charSource += numbers
	}

	if includeSpecial {
		charSource += specialSymbols
	}

	if useLetters {
		charSource += letters
	}

	for i := 0; i < length; i++ {
		randNum := rand.Intn(len(charSource))
		password = append(password, charSource[randNum])
	}

	return types.ToString(password)
}
