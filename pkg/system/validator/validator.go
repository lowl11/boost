package validator

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/lowl11/boost/errors"
	"net/http"
)

const (
	errorType  = "ValidateModel"
	contextKey = "validate"
)

type Validator struct {
	*baseValidator.Validate
	turnOff bool
}

func New() (*Validator, error) {
	base := baseValidator.New()

	if err := base.RegisterValidation("uuid", validateUUID); err != nil {
		return nil, err // TODO: need implement BoostError?
	}

	if err := base.RegisterValidation("undefined", validateUndefined); err != nil {
		return nil, err
	}

	return &Validator{
		Validate: base,
	}, nil
}

func (validator *Validator) TurnOff() *Validator {
	validator.turnOff = true
	return validator
}

func (validator *Validator) Struct(object any) error {
	if validator.turnOff {
		return nil
	}

	validateError := validator.Validate.Struct(object)
	if validateError == nil {
		return nil
	}

	err := errors.
		New("Model validation error").
		SetType(errorType).
		SetHttpCode(http.StatusUnprocessableEntity)

	validationErrors, ok := validateError.(baseValidator.ValidationErrors)
	if !ok {
		return err.SetError(validateError)
	}

	if len(validationErrors) == 0 {
		return nil
	}

	validations := make([]string, 0, len(validationErrors))
	for _, validationError := range validationErrors {
		validations = append(validations, validationError.Error())
	}

	return err.AddContext("validations", validations)
}

func (validator *Validator) Var(variable any, tag string) error {
	if validator.turnOff {
		return nil
	}

	err := errors.
		New("Variable validation error").
		SetType(errorType).
		SetHttpCode(http.StatusUnprocessableEntity)

	validateError := validator.Validate.Var(variable, tag)
	if err == nil {
		return err
	}

	validationErrors, ok := validateError.(baseValidator.ValidationErrors)
	if !ok {
		return err.AddContext(contextKey, validateError.Error())
	}

	if len(validationErrors) == 0 {
		return nil
	}

	validations := make([]string, 0, len(validationErrors))
	for _, validationError := range validationErrors {
		validations = append(validations, validationError.Error())
	}

	return err.AddContext(contextKey, validations)
}
