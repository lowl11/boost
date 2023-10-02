package jwt_service

import (
	"github.com/lowl11/boost/pkg/system/types"
)

type JWT struct {
	key []byte
}

func New(key string) *JWT {
	return &JWT{
		key: types.ToBytes(key),
	}
}
