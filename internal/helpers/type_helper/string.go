package type_helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func ToString(anyValue any, memory bool) string {
	if anyValue == nil {
		return ""
	}

	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	if bytesBuffer, ok := anyValue.([]byte); ok {
		return BytesToString(bytesBuffer)
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.String:
		return anyValue.(string)
	case reflect.Bool:
		return strconv.FormatBool(anyValue.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		valueInBytes, err := json.Marshal(anyValue)
		if err != nil {
			return ""
		}
		return string(valueInBytes)
	case reflect.Ptr:
		if memory {
			return fmt.Sprintf("%v", value)
		}

		return ToString(value.Elem().Interface(), true)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func BytesToString(buffer []byte) string {
	return *(*string)(unsafe.Pointer(&buffer))
}

func GetString(setValue string) string {
	if setValue == "" {
		return ""
	}

	return setValue
}

func StringFromError(err error, defaultMessage string) string {
	if err == nil {
		return defaultMessage
	}

	return err.Error()
}

func StringToBool(value string) bool {
	return value == "true"
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}
