package rmq_connection

import amqp "github.com/rabbitmq/amqp091-go"

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
