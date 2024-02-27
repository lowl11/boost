package boost

import (
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/pkg/web/auth/jwt"
)

func RegisterJWT(jwtKey string) {
	di.Register[jwt.JWT](func() *jwt.JWT {
		return jwt.New(jwtKey)
	})
}
