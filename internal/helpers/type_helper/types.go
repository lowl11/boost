package type_helper

import "reflect"

var (
	_primitives = []reflect.Kind{
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String,
		reflect.Array,
		reflect.Slice,
		reflect.Map,
	}
)

func IsPrimType(t reflect.Type) bool {
	for _, prim := range _primitives {
		if prim == t.Kind() {
			return true
		}
	}

	return false
}

func UnwrapType(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Ptr {
		return t
	}

	elem := t.Elem()
	for {
		if elem.Kind() != reflect.Ptr {
			break
		}

		elem = elem.Elem()
	}

	return elem
}

func UnwrapValue(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}

	elem := v.Elem()
	for {
		if elem.Kind() != reflect.Ptr {
			break
		}

		elem = elem.Elem()
	}

	return elem
}
