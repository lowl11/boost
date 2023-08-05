package rmq_service

import (
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	url string

	connection *amqp.Connection
	channel    *amqp.Channel
}

func New(url string) (*Service, error) {
	connection, err := rmq.NewConnection(url)
	if err != nil {
		return nil, err
	}

	channel, err := rmq.NewChannel(connection)
	if err != nil {
		return nil, err
	}

	return &Service{
		url: url,

		connection: connection,
		channel:    channel,
	}, nil
}
