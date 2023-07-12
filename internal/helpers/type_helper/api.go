package type_helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func BytesToString(buffer []byte) string {
	return *(*string)(unsafe.Pointer(&buffer))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  len(s),
		Cap:  len(s),
	}))
}

func ToString(anyValue any) string {
	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
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
		return ToString(value.Elem().Interface())
	default:
		return fmt.Sprintf("%v", value)
	}
}

func IsPrimitive(value any) bool {
	return isInteger(value) || isString(value) || isBool(value) || isFloat(value)
}

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

func isBool(value any) bool {
	_, ok := value.(bool)
	return ok
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}
