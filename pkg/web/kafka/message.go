package kafka

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost/pkg/io/types"
)

type message struct {
	topic     string
	partition int32
	offset    int64
	key       []byte
	value     []byte
	headers   map[string]string
}

func NewMessage(topic string, key, value []byte) Message {
	return &message{
		topic: topic,
		key:   key,
		value: value,
	}
}

func (msg *message) Topic() string {
	return msg.topic
}

func (msg *message) SetTopic(topic string) Message {
	msg.topic = topic
	return msg
}

func (msg *message) Partition() int32 {
	return msg.partition
}

func (msg *message) SetPartition(partition int32) Message {
	msg.partition = partition
	return msg
}

func (msg *message) Offset() int64 {
	return msg.offset
}

func (msg *message) SetOffset(offset int64) Message {
	msg.offset = offset
	return msg
}

func (msg *message) Key() []byte {
	return msg.key
}

func (msg *message) SetKey(key []byte) Message {
	msg.key = key
	return msg
}

func (msg *message) Value() []byte {
	return msg.value
}

func (msg *message) SetValue(value []byte) Message {
	msg.value = value
	return msg
}

func (msg *message) Headers() map[string]string {
	return msg.headers
}

func (msg *message) SetHeaders(headers map[string]string) Message {
	msg.headers = headers
	return msg
}

func messageFromConsumer(msg *sarama.ConsumerMessage) Message {
	headers := make(map[string]string, len(msg.Headers))
	for _, header := range msg.Headers {
		headers[types.String(header.Key)] = types.String(header.Value)
	}

	return &message{
		topic:     msg.Topic,
		partition: msg.Partition,
		offset:    msg.Offset,
		key:       msg.Key,
		value:     msg.Value,
		headers:   headers,
	}
}

func messageToProducer(msg Message) *sarama.ProducerMessage {
	headers := make([]sarama.RecordHeader, 0, len(msg.Headers()))
	for key, value := range msg.Headers() {
		headers = append(headers, sarama.RecordHeader{
			Key:   types.ToBytes(key),
			Value: types.ToBytes(value),
		})
	}

	return &sarama.ProducerMessage{
		Topic:   msg.Topic(),
		Key:     sarama.ByteEncoder(msg.Key()),
		Value:   sarama.ByteEncoder(msg.Value()),
		Headers: headers,
	}
}
