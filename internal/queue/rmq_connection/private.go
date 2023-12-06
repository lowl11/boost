package rmq_connection

import (
	"github.com/lowl11/boost/internal/queue/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (connection *Connection) getDispatcherChannel() (*amqp.Channel, error) {
	connection.dispatcherMutex.Lock()
	defer connection.dispatcherMutex.Unlock()

	if err := connection.reConnect(); err != nil {
		return nil, err
	}

	if connection.dispatcherChannel != nil && !connection.dispatcherChannel.IsClosed() {
		return connection.dispatcherChannel, nil
	}

	channel, err := connection.connection.Channel()
	if err != nil {
		return nil, err
	}

	connection.dispatcherChannel = channel
	return channel, nil
}

func (connection *Connection) getListenerChannel() (*amqp.Channel, error) {
	connection.listenerMutex.Lock()
	defer connection.listenerMutex.Unlock()

	if err := connection.reConnect(); err != nil {
		return nil, err
	}

	if connection.listenerChannel != nil && !connection.listenerChannel.IsClosed() {
		return connection.listenerChannel, nil
	}

	channel, err := connection.connection.Channel()
	if err != nil {
		return nil, err
	}

	connection.listenerChannel = channel
	return channel, nil
}

func (connection *Connection) reConnect() error {
	if !connection.connection.IsClosed() {
		return nil
	}

	var err error
	connection.connection, err = rmq.NewConnection(connection.connectionString)
	if err != nil {
		return err
	}

	return nil
}
