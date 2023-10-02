package types

import (
	"github.com/lowl11/boost/data/interfaces"
)

type HandlerFunc func(ctx interfaces.Context) error
