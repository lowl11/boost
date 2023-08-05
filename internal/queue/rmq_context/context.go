package rmq_context

import amqp "github.com/rabbitmq/amqp091-go"

type Context struct {
	body []byte
}

func New(message *amqp.Delivery) *Context {
	return &Context{
		body: message.Body,
	}
}
