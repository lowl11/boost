package type_helper

import (
	"reflect"
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
