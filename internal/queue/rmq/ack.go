package rmq

import amqp "github.com/rabbitmq/amqp091-go"

func Ack(channel *amqp.Channel, deliveryTag uint64, multiple bool) error {
	return channel.Ack(deliveryTag, multiple)
}
