package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/async"
	"github.com/lowl11/boost/pkg/io/exception"
	"github.com/lowl11/boost/pkg/system/cancel"
	"sync"
)

type consumer struct {
	ctx      context.Context
	cfg      *Config
	consumer sarama.Consumer
}

func NewConsumer(ctx context.Context, config *Config) (Consumer, error) {
	kafkaConsumer, err := sarama.NewConsumer(config.hosts, config.saramaConfig())
	if err != nil {
		return nil, err
	}

	return &consumer{
		ctx:      ctx,
		cfg:      config,
		consumer: kafkaConsumer,
	}, nil
}

func (c *consumer) StartListeningAsync(topic string, handler Handler) {
	async.Run(c.ctx, func(ctx context.Context) error {
		if err := c.StartListening(topic, handler); err != nil {
			return err
		}

		return nil
	})
}

func (c *consumer) StartListening(topic string, handler Handler) error {
	partitions, err := c.consumer.Partitions(topic)
	if err != nil {
		return errors.
			New("Get partitions error").
			SetType("Kafka_GetPartitions").
			SetError(err).
			AddContext("topic", topic)
	}

	goroutines := make([]*sync.WaitGroup, len(partitions))
	for i := 0; i < len(partitions); i++ {
		goroutines[i] = &sync.WaitGroup{}
		goroutines[i].Add(1)
	}

	for partition := 0; partition < len(goroutines); partition++ {
		go c.handleConsumerFunc(goroutines[partition], topic, int32(partition), handler)
	}

	<-c.ctx.Done()
	return c.consumer.Close()
}

func (c *consumer) handleConsumerFunc(wg *sync.WaitGroup, topic string, partitionNum int32, handler Handler) {
	defer wg.Done()

	partConsumer, err := c.consumer.ConsumePartition(
		topic,
		partitionNum,
		c.cfg.saramaConfig().Consumer.Offsets.Initial,
	)
	if err != nil {
		log.Error(errors.
			New("Consumer partition error").
			SetType("Kafka_ConsumePartition").
			SetError(err).
			SetContext(map[string]any{
				"topic":     topic,
				"partition": partitionNum,
				"offset":    c.cfg.saramaConfig().Consumer.Offsets.Initial,
			}))
		return
	}

	errorHandler := c.cfg.errorHandler
	go func() {
		for msg := range partConsumer.Messages() {
			if err = exception.Try(func() error {
				return handler(messageFromConsumer(msg))
			}); err != nil {
				if errorHandler != nil {
					errorHandler(err)
				} else {
					log.Error("Kafka consume message error:", err)
				}
			}
		}
	}()

	cancel.Get().Add()
	defer cancel.Get().Done()

	for {
		select {
		case kafkaError := <-partConsumer.Errors():
			log.Error("Kafka consumer error:", kafkaError.Error(), ". Partition:", kafkaError.Partition)
			return
		case <-c.ctx.Done():
			log.Info("Kafka consumer closed partition #", partitionNum+1)
			return
		}
	}
}

func (c *consumer) Close() error {
	return c.consumer.Close()
}
