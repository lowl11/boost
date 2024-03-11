package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/async"
	"github.com/lowl11/boost/pkg/system/cancel"
)

type consumerGroup struct {
	ctx      context.Context
	cfg      *Config
	consumer sarama.ConsumerGroup
}

func NewConsumerGroup(ctx context.Context, cfg *Config, group string) (ConsumerGroup, error) {
	kafkaConsumer, err := sarama.NewConsumerGroup(cfg.hosts, group, cfg.saramaConfig())
	if err != nil {
		return nil, err
	}

	return &consumerGroup{
		ctx:      ctx,
		cfg:      cfg,
		consumer: kafkaConsumer,
	}, nil
}

func (c *consumerGroup) StartListeningAsync(topic string, groupHandler sarama.ConsumerGroupHandler) {
	async.Run(c.ctx, func(ctx context.Context) error {
		if err := c.StartListening(topic, groupHandler); err != nil {
			return err
		}

		return nil
	})
}

func (c *consumerGroup) StartListening(topic string, groupHandler sarama.ConsumerGroupHandler) error {
	go func() {
		if err := c.consumer.Consume(c.ctx, []string{topic}, groupHandler); err != nil {
			log.Fatal("Start consuming group error:", err)
		}
	}()

	cancel.Get().Add()
	defer cancel.Get().Done()

	for {
		select {
		case <-c.ctx.Done():
			log.Info("Kafka consumer group closed")
			return nil
		}
	}
}

func (c *consumerGroup) Close() error {
	return c.consumer.Close()
}
