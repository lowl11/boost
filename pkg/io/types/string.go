package types

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lowl11/flex"
	"reflect"
	"strconv"
	"unsafe"
)

func String(value any) string {
	return toString(value, false)
}

func toString(anyValue any, memory bool) string {
	if anyValue == nil {
		return ""
	}

	// already string type
	if stringValue, isStr := anyValue.(string); isStr {
		return stringValue
	}

	// try cast to error
	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	// try cast to bytes
	if bytesBuffer, ok := anyValue.([]byte); ok {
		return BytesToString(bytesBuffer)
	}

	// try get Stringer interface
	if stringer, ok := anyValue.(fmt.Stringer); ok {
		return stringer.String()
	}

	// try cast uuid
	if flex.Type(reflect.TypeOf(anyValue)).IsUUID() {
		uuidValue, ok := anyValue.(uuid.UUID)
		if ok {
			return uuidValue.String()
		}

		uuidPtrValue, ok := anyValue.(*uuid.UUID)
		if ok {
			return uuidPtrValue.String()
		}
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

		return toString(value.Elem().Interface(), true)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func BytesToString(buffer []byte) string {
	return *(*string)(unsafe.Pointer(&buffer))
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}
