package rmq_connection

import amqp "github.com/rabbitmq/amqp091-go"

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
