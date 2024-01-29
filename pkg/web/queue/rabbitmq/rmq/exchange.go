package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type exchangeConfig struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

func defaultExchangeConfig() exchangeConfig {
	return exchangeConfig{}
}

func newExchange(channel *amqp.Channel, exchangeName, exchangeType string, cfg ...exchangeConfig) error {
	var config exchangeConfig
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
