package exception

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/internal/boosties/exception/try"
)

func Try(tryFunc func()) boost.Try {
	return try.New(tryFunc)
}
