package validator

import (
	baseValidator "github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

func validateUUID(fl baseValidator.FieldLevel) (isValid bool) {
	switch val := fl.Field().Interface().(type) {
	case string:
		isValid = uuid.FromStringOrNil(val) != uuid.Nil
	case *string:
		isValid = *val == "" || uuid.FromStringOrNil(*val) != uuid.Nil
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
