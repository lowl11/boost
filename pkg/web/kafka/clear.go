package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

func ClearTopic(ctx context.Context, cfg *Config, topic string, partitions int32) error {
	admin, err := sarama.NewClusterAdmin(cfg.hosts, cfg.saramaConfig())
	if err != nil {
		return err
	}

	if err = admin.DeleteTopic(topic); err != nil && err.Error() != sarama.ErrUnknownTopicOrPartition.Error() {
		return err
	}

	return admin.CreateTopic(
		topic,
		&sarama.TopicDetail{
			NumPartitions:     partitions,
			ReplicationFactor: -1,
		},
		false,
	)
}
