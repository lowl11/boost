package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type ExchangeConfig struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

func defaultExchangeConfig() ExchangeConfig {
	return ExchangeConfig{}
}

func NewExchange(channel *amqp.Channel, exchangeName, exchangeType string, cfg ...ExchangeConfig) error {
	var config ExchangeConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultExchangeConfig()
	}

	return channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		config.Durable,
		config.AutoDelete,
		config.Internal,
		config.NoWait,
		config.Args,
	)
}
