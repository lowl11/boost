package types

import (
	"github.com/lowl11/boost/data/interfaces"
)

type MiddlewareFunc func(ctx interfaces.Context) error
