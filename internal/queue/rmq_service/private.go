package rmq_service

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/internal/queue/rmq"
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
	service.channel = channel

	return nil
}
