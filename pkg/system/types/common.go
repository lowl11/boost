package types

func IsPrimitive(value any) bool {
	return isInteger(value) || isString(value) || isBool(value) || isFloat(value)
}

func isBool(value any) bool {
	_, ok := value.(bool)
	return ok
}
