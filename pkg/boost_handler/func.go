package boost_handler

import "github.com/lowl11/boost/pkg/boost_context"

type HandlerFunc func(ctx boost_context.Context) error
