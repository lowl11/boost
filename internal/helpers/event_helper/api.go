package event_helper

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"reflect"
)

func NameOfEvent(event any) (string, error) {
	// func getType(myvar interface{}) string {
	//    if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
	//        return "*" + t.Elem().Name()
	//    } else {
	//        return t.Name()
	//    }
	//}

	reflectValue := reflect.TypeOf(event)
	if reflectValue.Kind() != reflect.Struct {
		return "", errors.New("given argument is not struct{}")
	}

	return reflectValue.Name(), nil
}
