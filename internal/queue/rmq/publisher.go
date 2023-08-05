package rmq

import (
	"context"
	"github.com/lowl11/boost/pkg/enums/content_types"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishConfig struct {
	Exchange    string
	Mandatory   bool
	Immediate   bool
	ContentType string
}

func defaultPublishConfig() PublishConfig {
	return PublishConfig{
		ContentType: content_types.JSON,
	}
}

func Publish(ctx context.Context, channel *amqp.Channel, queueName string, body []byte, cfg ...PublishConfig) error {
	var config PublishConfig
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
