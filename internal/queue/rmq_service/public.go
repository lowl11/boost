package rmq_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (service *Service) Publish(ctx context.Context, exchangeName, eventName string, event any) error {
	eventInBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	dispatcherChannel, err := service.getDispatcherChannel()
	if err != nil {
		return err
	}

	// first try
	if err = rmq.Publish(ctx, dispatcherChannel, eventName, eventInBytes, rmq.PublishConfig{
		Exchange: exchangeName,
	}); err != nil {
		// if connection was closed
		if errors.Is(err, amqp.ErrClosed) {
			// reconnect to RMQ
			if err = service.reConnect(); err != nil {
				return err
			}

			// second try
			if err = rmq.Publish(ctx, dispatcherChannel, eventName, eventInBytes, rmq.PublishConfig{
				Exchange: exchangeName,
			}); err != nil {
				return err
			}

			return nil
		}

		return err
	}

	return nil
}

func (service *Service) Ack(deliveryTag uint64) error {
	listenerChannel, err := service.getListenerChannel()
	if err != nil {
		return err
	}

	return rmq.Ack(listenerChannel, deliveryTag, false)
}

func (service *Service) NewExchange(exchangeName, exchangeType string) error {
	listenerChannel, err := service.getListenerChannel()
	if err != nil {
		return err
	}

	return rmq.NewExchange(listenerChannel, exchangeName, exchangeType)
}

func (service *Service) NewQueue(queueName string) (*amqp.Queue, error) {
	listenerChannel, err := service.getListenerChannel()
	if err != nil {
		return nil, err
	}

	return rmq.NewQueue(listenerChannel, queueName)
}

func (service *Service) Bind(exchangeName, queueName string) error {
	listenerChannel, err := service.getListenerChannel()
	if err != nil {
		return err
	}

	return rmq.Bind(listenerChannel, exchangeName, queueName)
}

func (service *Service) Consume(queueName string) (<-chan amqp.Delivery, error) {
	listenerChannel, err := service.getListenerChannel()
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
	if service.connection == nil {
		return nil
	}

	if service.dispatcherChannel != nil {
		if err := service.dispatcherChannel.Close(); err != nil {
			return err
		}
	}

	if service.listenerChannel != nil {
		if err := service.listenerChannel.Close(); err != nil {
			return err
		}
	}

	return service.connection.Close()
}
