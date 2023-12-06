package rmq_service

import (
	"context"
	"encoding/json"
	"github.com/lowl11/boost/internal/queue/rmq"
	"github.com/lowl11/boost/internal/queue/rmq_connection"
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
	if err = rmq.Publish(ctx, dispatcherChannel, eventName, eventInBytes, rmq.PublishConfig{
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

	return rmq.Ack(listenerChannel, deliveryTag, false)
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

	return rmq.NewExchange(listenerChannel, exchangeName, exchangeType)
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

	return rmq.NewQueue(listenerChannel, queueName)
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

	return rmq.Bind(listenerChannel, exchangeName, queueName)
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

	messages, err := rmq.Consume(listenerChannel, queueName)
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
