package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

type ClearConfig struct {
	// NumPartitions contains the number of partitions to create in the topic, or
	// -1 if we are either specifying a manual partition assignment or using the
	// default partitions.
	NumPartitions int32
	// ReplicationFactor contains the number of replicas to create for each
	// partition in the topic, or -1 if we are either specifying a manual
	// partition assignment or using the default replication factor.
	ReplicationFactor int16
	// ReplicaAssignment contains the manual partition assignment, or the empty
	// array if we are using automatic assignment.
	ReplicaAssignment map[int32][]int32
	// ConfigEntries contains the custom topic configurations to set.
	ConfigEntries map[string]*string
}

func defaultClearConfig(cfg *ClearConfig) ClearConfig {
	var partitions int32
	if cfg.NumPartitions > 0 {
		partitions = cfg.NumPartitions
	} else {
		partitions = 1
	}

	var replicationFactor int16
	if cfg.ReplicationFactor != 0 {
		replicationFactor = cfg.ReplicationFactor
	} else {
		replicationFactor = -1
	}

	return ClearConfig{
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
		ReplicaAssignment: cfg.ReplicaAssignment,
		ConfigEntries:     cfg.ConfigEntries,
	}
}

func ClearTopic(ctx context.Context, cfg *Config, topic string, clearCfg ClearConfig) error {
	clearCfg = defaultClearConfig(&clearCfg)

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
			NumPartitions:     clearCfg.NumPartitions,
			ReplicationFactor: clearCfg.ReplicationFactor,
			ReplicaAssignment: clearCfg.ReplicaAssignment,
			ConfigEntries:     clearCfg.ConfigEntries,
		},
		false,
	)
}
