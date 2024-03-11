package interfaces

import (
	"context"
)

type Dispatcher interface {
	Dispatch(ctx context.Context, event any) error
}

type Listener interface {
	Run(amqpConnectionURL string) error
	Bind(event any, action func(ctx EventContext) error)
	EventsCount() int
	Close() error
}

type EventContext interface {
	Body() []byte
	Parse(object any) error
	Context() context.Context
}
