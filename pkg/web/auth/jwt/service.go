package jwt

import (
	"github.com/lowl11/boost/pkg/io/types"
)

// JWT is a service which generates & parse JWT tokens
type JWT struct {
	key []byte
}

// New creates new JWT service instance
func New(key string) *JWT {
	return &JWT{
		key: types.ToBytes(key),
	}
}
