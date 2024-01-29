package rmq

import (
	"context"
	"github.com/lowl11/boost/data/enums/content_types"
	amqp "github.com/rabbitmq/amqp091-go"
)

type publishConfig struct {
	Exchange    string
	Mandatory   bool
	Immediate   bool
	ContentType string
}

func defaultPublishConfig() publishConfig {
	return publishConfig{
		ContentType: content_types.JSON,
	}
}

func publish(ctx context.Context, channel *amqp.Channel, queueName string, body []byte, cfg ...publishConfig) error {
	var config publishConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultPublishConfig()
	}

	err := channel.PublishWithContext(ctx,
		config.Exchange,
		queueName,
		config.Mandatory,
		config.Immediate,
		amqp.Publishing{
			ContentType: config.ContentType,
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}
