package errors

import (
	"github.com/lowl11/boost/data/interfaces"
	native "github.com/pkg/errors"
)

// New creates new Boost Error object with given message
func New(message string) interfaces.Error {
	return newError(message)
}

func Parse(response []byte) (interfaces.Error, bool) {
	return parseCustom(response)
}

func IsType(err error, errorType string) bool {
	boostError, ok := err.(interfaces.Error)
	if !ok {
		return false
	}

	return boostError.Type() == errorType
}

func Is(err1, err2 error) bool {
	boostError1, isBoost := err1.(interfaces.Error)
	if !isBoost {
		return native.Is(err1, err2)
	}

	boostError2, isBoost := err2.(interfaces.Error)
	if !isBoost {
		return native.Is(err1, err2)
	}

	return boostError1.Is(boostError2)
}
