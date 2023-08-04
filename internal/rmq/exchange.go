package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type ExchangeConfig struct {
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

func NewExchange(channel *amqp.Channel, exchangeName string, cfg ExchangeConfig) error {
	return channel.ExchangeDeclare(
		exchangeName,
		cfg.Kind,
		cfg.Durable,
		cfg.AutoDelete,
		cfg.Internal,
		cfg.NoWait,
		cfg.Args,
	)
}
