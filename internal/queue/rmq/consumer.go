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

func defaultConsumerConfig() ConsumerConfig {
	return ConsumerConfig{}
}

func Consume(channel *amqp.Channel, queueName string, cfg ...ConsumerConfig) (MessagesQueue, error) {
	var config ConsumerConfig
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
