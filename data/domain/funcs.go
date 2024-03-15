package domain

import "github.com/lowl11/boost/data/interfaces"

type HandlerFunc func(ctx interfaces.Context) error
type MiddlewareFunc func(ctx interfaces.Context) error
type ListenerAction func(event interfaces.EventContext) error
type PanicHandler = func(err error)
