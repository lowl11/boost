package rmq_connection

import (
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

	connection, err := newConnection(_connectionString)
	if err != nil {
		return nil, err
	}

	instance = &Connection{
		connectionString: _connectionString,
		connection:       connection,
	}
	return instance, nil
}

func (connection *Connection) GetDispatcherChannel() (*amqp.Channel, error) {
	return connection.getDispatcherChannel()
}

func (connection *Connection) GetListenerChannel() (*amqp.Channel, error) {
	return connection.getListenerChannel()
}

func (connection *Connection) Close() error {
	if connection.dispatcherChannel != nil {
		if !connection.dispatcherChannel.IsClosed() {
			_ = connection.dispatcherChannel.Close()
		}
	}

	if connection.listenerChannel != nil {
		if !connection.listenerChannel.IsClosed() {
			_ = connection.listenerChannel.Close()
		}
	}

	if connection.connection.IsClosed() {
		return nil
	}

	return connection.connection.Close()
}

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
	connection.connection, err = newConnection(connection.connectionString)
	if err != nil {
		return err
	}

	return nil
}

func newConnection(url string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func newChannel(connection *amqp.Connection) (*amqp.Channel, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
