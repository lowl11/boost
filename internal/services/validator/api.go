package validator

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

func (validator *Validator) Struct(object any) error {
	err := errors.
		New("Model validation error").
		SetType("ValidateModel").
		SetHttpCode(http.StatusUnprocessableEntity)

	validateError := validator.Validate.Struct(object)
	validationErrors, ok := validateError.(baseValidator.ValidationErrors)
	if !ok {
		return err.SetError(err)
	}

	validations := make([]string, 0, len(validationErrors))
	for _, validationError := range validationErrors {
		validations = append(validations, validationError.Error())
	}

	return err.AddContext("validation", validations)
}
