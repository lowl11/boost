package event_helper

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"reflect"
)

func NameOfEvent(event any) (string, error) {
	reflectValue := reflect.TypeOf(event)
	if reflectValue.Kind() != reflect.Struct {
		return "", errors.New("given argument is not struct{}")
	}

	return reflectValue.Name(), nil
}
