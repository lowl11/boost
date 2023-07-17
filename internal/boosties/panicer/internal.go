package panicer

import "errors"

func fromAny(err any) error {
	if err == nil {
		return nil
	}

	switch err.(type) {
	case string:
		return errors.New(err.(string))
	case error:
		return errors.New(err.(error).Error())
	}

	return nil
}
