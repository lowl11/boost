package types

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

func ToBytes(anyValue any, memory bool) []byte {
	if anyValue == nil {
		return nil
	}

	valueType := reflect.TypeOf(anyValue)

	switch valueType.Kind() {
	case reflect.String:
		return StringToBytes(anyValue.(string))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Bool:
		return StringToBytes(toString(anyValue, false))
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		marshalled, err := json.Marshal(anyValue)
		if err != nil {
			return nil
		}

		return marshalled
	case reflect.Ptr:
		if memory {
			return nil
		}

		return ToBytes(reflect.ValueOf(anyValue).Interface(), true)
	default:
		return StringToBytes(toString(anyValue, false))
	}
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  len(s),
		Cap:  len(s),
	}))
}

func StringSliceToBytes(slice []string) ([]byte, error) {
	marshalled, err := json.Marshal(slice)
	if err != nil {
		return nil, err
	}

	return marshalled, nil
}
