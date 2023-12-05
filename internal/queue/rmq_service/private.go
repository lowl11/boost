package rmq_service

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (service *Service) reConnect(memory ...bool) error {
	var returnError bool

	if len(memory) > 0 {
		if memory[0] {
			returnError = true
			return errors.
				New("Connect to RabbitMQ error").
				SetType("RMQ_ConnectError")
		}
	}

	connection, err := rmq.NewConnection(service.url)
	if err != nil {
		if returnError {
			return errors.
				New("Connect to RabbitMQ error").
				SetType("RMQ_ConnectError")
		}

		return service.reConnect(true)
	}

	channel, err := rmq.NewChannel(connection)
	if err != nil {
		if returnError {
			return errors.
				New("Connect to RabbitMQ error").
				SetType("RMQ_ConnectError")
		}

		return service.reConnect(true)
	}

	service.connection = connection
	service.dispatcherChannel = channel

	return nil
}

func (service *Service) getDispatcherChannel() (*amqp.Channel, error) {
	if service.dispatcherChannel != nil {
		return service.dispatcherChannel, nil
	}

	dispatcherChannel, err := service.connection.Channel()
	if err != nil {
		return nil, err
	}
	service.dispatcherChannel = dispatcherChannel
	return dispatcherChannel, nil
}

func (service *Service) getListenerChannel() (*amqp.Channel, error) {
	if service.listenerChannel != nil {
		return service.listenerChannel, nil
	}

	listenerChannel, err := service.connection.Channel()
	if err != nil {
		return nil, err
	}
	service.listenerChannel = listenerChannel
	return listenerChannel, nil
}
