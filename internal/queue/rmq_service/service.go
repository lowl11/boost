package rmq_service

import (
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	url string

	connection        *amqp.Connection
	listenerChannel   *amqp.Channel
	dispatcherChannel *amqp.Channel
}

func New(url string) (*Service, error) {
	connection, err := rmq.NewConnection(url)
	if err != nil {
		return nil, err
	}

	return &Service{
		url:        url,
		connection: connection,
	}, nil
}
