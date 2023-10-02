package types

import "github.com/lowl11/boost/pkg/interfaces"

type MiddlewareFunc func(ctx interfaces.Context) error
