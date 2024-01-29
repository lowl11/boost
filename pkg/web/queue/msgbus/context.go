package msgbus

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type eventContext struct {
	body []byte
}

func newContext(message *amqp.Delivery) *eventContext {
	return &eventContext{
		body: message.Body,
	}
}

func (ctx *eventContext) Body() []byte {
	return ctx.body
}

func (ctx *eventContext) Parse(object any) error {
	return json.Unmarshal(ctx.body, &object)
}
