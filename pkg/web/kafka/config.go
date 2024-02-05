package kafka

import (
	"github.com/IBM/sarama"
	"time"
)

type Config struct {
	hosts         []string
	errorHandler  ErrorHandler
	username      string
	password      string
	saslMechanism string

	saramaCfg *sarama.Config

	opts []Option
}

func NewConfig(hosts []string) *Config {
	return &Config{
		hosts: hosts,
		opts:  make([]Option, 0),
	}
}

func (config *Config) saramaConfig() *sarama.Config {
	if config.saramaCfg != nil {
		return config.saramaCfg
	}

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true

	if config.username != "" {
		saramaConfig.Net.SASL.Enable = true
		saramaConfig.Net.SASL.User = config.username
		saramaConfig.Net.SASL.Password = config.password
		saramaConfig.Net.SASL.Handshake = true
		if config.saslMechanism == "" {
			saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		} else {
			saramaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(config.saslMechanism)
		}
	}

	return saramaConfig
}

func (config *Config) SetHosts(hosts []string) *Config {
	config.hosts = hosts
	return config
}

func (config *Config) SetErrorHandler(handler ErrorHandler) *Config {
	config.errorHandler = handler
	return config
}

func (config *Config) WithAutocommit(interval time.Duration) *Config {
	config.opts = append(config.opts, func(config *sarama.Config) {
		config.Consumer.Offsets.AutoCommit.Enable = true
		config.Consumer.Offsets.AutoCommit.Interval = interval
	})
	return config
}

func (config *Config) With(optionFunc Option) *Config {
	config.opts = append(config.opts, optionFunc)
	return config
}
