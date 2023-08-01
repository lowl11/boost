package errors

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/interfaces"
)

// New creates new Boost Error object with given message
func New(message string) interfaces.Error {
	return errors.New(message)
}

func Parse(response []byte) (interfaces.Error, bool) {
	return errors.Parse(response)
}
