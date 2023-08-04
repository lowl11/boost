package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type MessagesQueue <-chan amqp.Delivery

type ConsumerConfig struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func Consume(channel *amqp.Channel, queue *amqp.Queue, cfg ConsumerConfig) (MessagesQueue, error) {
	messages, err := channel.Consume(
		queue.Name,
		cfg.Consumer,
		cfg.AutoAck,
		cfg.Exclusive,
		cfg.NoLocal,
		cfg.NoWait,
		cfg.Args,
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
