package rmq

import amqp "github.com/rabbitmq/amqp091-go"

func bind(channel *amqp.Channel, exchangeName, queueName string) error {
	if err := channel.QueueBind(
		queueName,
		queueName,
		exchangeName,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
