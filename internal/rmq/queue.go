package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type QueueConfig struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func NewQueue(channel *amqp.Channel, queueName string, cfg QueueConfig) (*amqp.Queue, error) {
	someQueue, err := channel.QueueDeclare(
		queueName,
		cfg.Durable,
		cfg.AutoDelete,
		cfg.Exclusive,
		cfg.NoWait,
		cfg.Args,
	)
	if err != nil {
		return nil, err
	}

	return &someQueue, nil
}
