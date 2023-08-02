package type_helper

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

func ToBytes(anyValue any) []byte {
	if anyValue == nil {
		return nil
	}

	valueType := reflect.TypeOf(anyValue)

	switch valueType.Kind() {
	case reflect.String:
		return StringToBytes(anyValue.(string))
	default:
		return StringToBytes(ToString(anyValue))
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
