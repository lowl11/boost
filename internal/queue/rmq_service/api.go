package rmq_service

import (
	"context"
	"encoding/json"
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (service *Service) Publish(ctx context.Context, exchangeName, eventName string, event any) error {
	eventInBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err = rmq.Publish(ctx, service.channel, eventName, eventInBytes, rmq.PublishConfig{
		Exchange: exchangeName,
	}); err != nil {
		return err
	}

	return nil
}

func (service *Service) Ack(deliveryTag uint64) error {
	return rmq.Ack(service.channel, deliveryTag, true)
}

func (service *Service) NewExchange(exchangeName, exchangeType string) error {
	return rmq.NewExchange(service.channel, exchangeName, exchangeType)
}

func (service *Service) NewQueue(queueName string) (*amqp.Queue, error) {
	return rmq.NewQueue(service.channel, queueName)
}

func (service *Service) Bind(exchangeName, queueName string) error {
	return rmq.Bind(service.channel, exchangeName, queueName)
}

func (service *Service) Consume(queueName string) (<-chan amqp.Delivery, error) {
	messages, err := rmq.Consume(service.channel, queueName)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (service *Service) Close() error {
	if service.connection == nil {
		return nil
	}

	if err := service.channel.Close(); err != nil {
		return err
	}

	return service.connection.Close()
}
