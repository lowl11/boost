package kafka

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost/errors"
)

type syncProducer struct {
	producer sarama.SyncProducer
	cfg      *Config
}

func NewSyncProducer(config *Config) (SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(config.hosts, config.saramaConfig())
	if err != nil {
		return nil, errors.
			New("Create Kafka Sync Producer error").
			SetType("Kafka_CreateSyncProducer").
			SetError(err)
	}

	return &syncProducer{
		producer: producer,
		cfg:      config,
	}, nil
}

func (p *syncProducer) Produce(messages ...Message) error {
	if len(messages) == 0 {
		return nil
	}

	saramaMessages := make([]*sarama.ProducerMessage, 0, len(messages))
	for _, msg := range messages {
		saramaMessages = append(saramaMessages, messageToProducer(msg))
	}

	if err := p.producer.SendMessages(saramaMessages); err != nil {
		return errors.
			New("Producer (sync) message error").
			SetType("Kafka_SyncProduce").
			SetError(err)
	}

	return nil
}
