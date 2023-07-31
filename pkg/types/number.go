package types

import (
	"github.com/lowl11/boost/internal/helpers/type_helper"
)

func ToInt(anyValue any) int {
	return type_helper.ToInt(anyValue)
}

func ToFloat32(anyValue any) float32 {
	return type_helper.ToFloat32(anyValue)
}

func ToFloat64(anyValue any) float64 {
	return type_helper.ToFloat64(anyValue)
}
