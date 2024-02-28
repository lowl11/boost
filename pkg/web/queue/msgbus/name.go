package msgbus

import (
	"github.com/lowl11/boost/errors"
	"reflect"
)

func nameOfEvent(event any) (string, error) {
	reflectValue := reflect.TypeOf(event)
	if reflectValue.Kind() != reflect.Struct {
		return "", errors.New("given argument is not struct{}")
	}

	return reflectValue.Name(), nil
}
