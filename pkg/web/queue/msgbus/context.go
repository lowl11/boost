package msgbus

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type eventContext struct {
	body []byte
	ctx  context.Context
}

func newContext(ctx context.Context, message *amqp.Delivery) *eventContext {
	return &eventContext{
		ctx:  ctx,
		body: message.Body,
	}
}

func (ctx *eventContext) Body() []byte {
	return ctx.body
}

func (ctx *eventContext) Parse(object any) error {
	return json.Unmarshal(ctx.body, &object)
}

func (ctx *eventContext) Context() context.Context {
	if ctx.ctx == nil {
		return context.Background()
	}

	return ctx.ctx
}
