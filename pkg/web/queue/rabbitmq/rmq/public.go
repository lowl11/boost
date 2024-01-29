package rmq

import (
	"context"
	"encoding/json"
	"github.com/lowl11/boost/pkg/web/queue/rabbitmq/rmq_connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (service *Service) Publish(ctx context.Context, exchangeName, eventName string, event any) error {
	eventInBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	connection, err := rmq_connection.Get()
	if err != nil {
		return err
	}

	dispatcherChannel, err := connection.GetDispatcherChannel()
	if err != nil {
		return err
	}

	// first try
	if err = publish(ctx, dispatcherChannel, eventName, eventInBytes, publishConfig{
		Exchange: exchangeName,
	}); err != nil {
		return err
	}

	return nil
}

func (service *Service) Ack(deliveryTag uint64) error {
	connection, err := rmq_connection.Get()
	if err != nil {
		return err
	}

	listenerChannel, err := connection.GetListenerChannel()
	if err != nil {
		return err
	}

	return ack(listenerChannel, deliveryTag, false)
}

func (service *Service) NewExchange(exchangeName, exchangeType string) error {
	connection, err := rmq_connection.Get()
	if err != nil {
		return err
	}

	listenerChannel, err := connection.GetListenerChannel()
	if err != nil {
		return err
	}

	return newExchange(listenerChannel, exchangeName, exchangeType)
}

func (service *Service) NewQueue(queueName string) (*amqp.Queue, error) {
	connection, err := rmq_connection.Get()
	if err != nil {
		return nil, err
	}

	listenerChannel, err := connection.GetListenerChannel()
	if err != nil {
		return nil, err
	}

	return newQueue(listenerChannel, queueName)
}

func (service *Service) Bind(exchangeName, queueName string) error {
	connection, err := rmq_connection.Get()
	if err != nil {
		return err
	}

	listenerChannel, err := connection.GetListenerChannel()
	if err != nil {
		return err
	}

	return bind(listenerChannel, exchangeName, queueName)
}

func (service *Service) Consume(queueName string) (<-chan amqp.Delivery, error) {
	connection, err := rmq_connection.Get()
	if err != nil {
		return nil, err
	}

	listenerChannel, err := connection.GetListenerChannel()
	if err != nil {
		return nil, err
	}

	messages, err := consume(listenerChannel, queueName)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (service *Service) Close() error {
	connection, err := rmq_connection.Get()
	if err != nil {
		return err
	}

	return connection.Close()
}
