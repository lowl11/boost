package types

import "github.com/lowl11/boost/internal/helpers/type_helper"

func ToString(value any) string {
	return type_helper.ToString(value, false)
}

func ToBytes(value any) []byte {
	return type_helper.ToBytes(value, false)
}

func StringToBool(value string) bool {
	return type_helper.StringToBool(value)
}
