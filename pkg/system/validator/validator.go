package validator

import baseValidator "github.com/go-playground/validator/v10"

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
