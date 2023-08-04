package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type QueueConfig struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func defaultQueueConfig() QueueConfig {
	return QueueConfig{}
}

func NewQueue(channel *amqp.Channel, queueName string, cfg ...QueueConfig) (*amqp.Queue, error) {
	var config QueueConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultQueueConfig()
	}

	queue, err := channel.QueueDeclare(
		queueName,
		config.Durable,
		config.AutoDelete,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}
