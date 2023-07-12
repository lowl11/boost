package errors

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/pkg/interfaces"
)

func New(message string) interfaces.Error {
	return errors.New(message)
}
