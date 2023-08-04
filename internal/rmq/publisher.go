package rmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishConfig struct {
	Exchange    string
	Mandatory   bool
	Immediate   bool
	ContentType string
}

func Publish(ctx context.Context, channel *amqp.Channel, queue *amqp.Queue, body []byte, cfg PublishConfig) error {
	err := channel.PublishWithContext(ctx,
		cfg.Exchange,
		queue.Name,
		cfg.Mandatory,
		cfg.Immediate,
		amqp.Publishing{
			ContentType: cfg.ContentType,
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}
