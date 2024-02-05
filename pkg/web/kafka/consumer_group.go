package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost/log"
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

	// todo: implement of close kafka consumer

	return &consumerGroup{
		ctx:      ctx,
		cfg:      cfg,
		consumer: kafkaConsumer,
	}, nil
}

func (c *consumerGroup) StartListeningAsync(topic string, groupHandler sarama.ConsumerGroupHandler) {
	if err := c.StartListening(topic, groupHandler); err != nil {
		log.Fatal("Start listening group async error:", err)
	}
}

func (c *consumerGroup) StartListening(topic string, groupHandler sarama.ConsumerGroupHandler) error {
	go func() {
		if err := c.consumer.Consume(c.ctx, []string{topic}, groupHandler); err != nil {
			log.Fatal("Start consuming group error:", err)
		}
	}()

	for {
		select {
		case <-c.ctx.Done():
			log.Info("Receiving group messages stopped by Context")
			return nil
		}
	}
}
