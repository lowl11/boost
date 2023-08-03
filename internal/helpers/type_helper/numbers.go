package type_helper

import (
	"reflect"
	"strconv"
)

func isInteger(value any) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		return true
	}
	return false
}

func isFloat(value any) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
		return true
	}
	return false
}

func ToInt(anyValue any) int {
	if anyValue == nil {
		return 0
	}

	if ptrValue, isPtr := anyValue.(*int); isPtr {
		return *ptrValue
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.String:
		stringValue := ToString(anyValue, false)
		if stringValue == "" {
			return 0
		}

		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return 0
		}

		return intValue
	case reflect.Float32, reflect.Float64:
		if floatValue, is32 := anyValue.(float32); is32 {
			return int(floatValue)
		}

		if floatValue, is64 := anyValue.(float64); is64 {
			return int(floatValue)
		}

		return 0
	}

	return 0
}

func ToFloat32(anyValue any) float32 {
	if anyValue == nil {
		return 0
	}

	if ptrValue, ok := anyValue.(*float32); ok {
		return *ptrValue
	}

	if floatValue, ok := anyValue.(float32); ok {
		return floatValue
	}

	if stringValue, ok := anyValue.(string); ok {
		floatValue, err := strconv.ParseFloat(stringValue, 32)
		if err != nil {
			return 0
		}

		return float32(floatValue)
	}

	return 0
}

func ToFloat64(anyValue any) float64 {
	if anyValue == nil {
		return 0
	}

	if ptrValue, isPtr := anyValue.(*float64); isPtr {
		return *ptrValue
	}

	if floatValue, ok := anyValue.(float64); ok {
		return floatValue
	}

	if stringValue, ok := anyValue.(string); ok {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return 0
		}

		return floatValue
	}

	return 0
}
