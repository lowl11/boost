package rmq_connection

import (
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type Connection struct {
	connectionString string

	connection        *amqp.Connection
	dispatcherChannel *amqp.Channel
	listenerChannel   *amqp.Channel

	connectionMutex sync.Mutex
	dispatcherMutex sync.Mutex
	listenerMutex   sync.Mutex
}

var instance *Connection
var _connectionString string

func SetURL(connectionString string) {
	if _connectionString != "" {
		return
	}

	_connectionString = connectionString
}

func Get() (*Connection, error) {
	if instance != nil {
		return instance, nil
	}

	connection, err := rmq.NewConnection(_connectionString)
	if err != nil {
		return nil, err
	}

	instance = &Connection{
		connectionString: _connectionString,
		connection:       connection,
	}
	return instance, nil
}
