package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type MessagesQueue <-chan amqp.Delivery

type consumerConfig struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func defaultConsumerConfig() consumerConfig {
	return consumerConfig{}
}

func consume(channel *amqp.Channel, queueName string, cfg ...consumerConfig) (MessagesQueue, error) {
	var config consumerConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultConsumerConfig()
	}

	messages, err := channel.Consume(
		queueName,
		config.Consumer,
		config.AutoAck,
		config.Exclusive,
		config.NoLocal,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
