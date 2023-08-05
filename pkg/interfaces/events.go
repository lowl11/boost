package interfaces

import (
	"context"
)

type Dispatcher interface {
	Dispatch(ctx context.Context, event any) error
}

type Listener interface {
	Run() error
	Bind(event any, action func(ctx EventContext) error)
}

type EventContext interface {
	Body() []byte
	Parse(object any) error
}
