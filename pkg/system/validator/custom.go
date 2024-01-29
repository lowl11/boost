package validator

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func validateUUID(fl baseValidator.FieldLevel) (isValid bool) {
	switch val := fl.Field().Interface().(type) {
	case string:
		_, err := uuid.Parse(val)
		if err != nil {
			return false
		}

		return true
	case *string:
		_, err := uuid.Parse(*val)
		if err != nil {
			return false
		}

		return true
	case uuid.UUID:
		isValid = val != uuid.Nil
	case *uuid.UUID:
		if fl.Field().IsNil() {
			return
		}

		isValid = *val != uuid.Nil
	}

	return
}

func validateUndefined(fl baseValidator.FieldLevel) (isValid bool) {
	const undefined = "undefined"
	switch val := fl.Field().Interface().(type) {
	case string:
		return val != undefined
	case *string:
		value := *val
		return value != undefined
	}

	return
}
