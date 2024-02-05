package kafka

import "github.com/IBM/sarama"

type Handler = func(message Message) error
type ErrorHandler = func(err error)
type Option func(config *sarama.Config)

type Consumer interface {
	StartListening(topic string, handler Handler) error
	StartListeningAsync(topic string, handler Handler)
}

type SyncProducer interface {
	Produce(messages ...Message) error
}

type Message interface {
	Topic() string
	SetTopic(topic string) Message

	Partition() int32
	SetPartition(partition int32) Message

	Offset() int64
	SetOffset(offset int64) Message

	Key() []byte
	SetKey(key []byte) Message

	Value() []byte
	SetValue(value []byte) Message

	Headers() map[string]string
	SetHeaders(map[string]string) Message
}
