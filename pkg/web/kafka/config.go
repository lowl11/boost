package kafka

import (
	"github.com/IBM/sarama"
	"time"
)

// Config is object necessary for running Kafka consumer & producer
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

	// default configs
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Consumer.Return.Errors = true

	// apply options
	for _, opt := range config.opts {
		opt(saramaConfig)
	}

	// default offset
	if saramaConfig.Consumer.Offsets.Initial == 0 {
		saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	return saramaConfig
}

// SetHosts set Kafka hosts
func (config *Config) SetHosts(hosts []string) *Config {
	config.hosts = hosts
	return config
}

// SetErrorHandler sets error handler for handler errors.
// For consumer - Handler function type error catch method.
func (config *Config) SetErrorHandler(handler ErrorHandler) *Config {
	config.errorHandler = handler
	return config
}

// WithAuth enables authorization for Kafka (via SASL)
func (config *Config) WithAuth(username, password string, mechanism ...string) *Config {
	config.opts = append(config.opts, func(config *sarama.Config) {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = username
		config.Net.SASL.Password = password
		config.Net.SASL.Handshake = true
		if len(mechanism) == 0 || mechanism[0] == "" {
			config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		} else {
			config.Net.SASL.Mechanism = sarama.SASLMechanism(mechanism[0])
		}
	})
	return config
}

// WithAutocommit fires autocommit mode with some interval.
// By default, interval is 1 second
func (config *Config) WithAutocommit(interval time.Duration) *Config {
	config.opts = append(config.opts, func(config *sarama.Config) {
		config.Consumer.Offsets.AutoCommit.Enable = true
		config.Consumer.Offsets.AutoCommit.Interval = interval
	})
	return config
}

// WithOffset sets start/initial offset for Consumer & ConsumerGroup
func (config *Config) WithOffset(offset int64) *Config {
	config.opts = append(config.opts, func(config *sarama.Config) {
		config.Consumer.Offsets.Initial = offset
	})
	return config
}

// With is dynamic option func.
// Developer make anything to Config object with this option.
func (config *Config) With(optionFunc Option) *Config {
	config.opts = append(config.opts, optionFunc)
	return config
}

// Copy copies Config object to new one but with the same data
func (config *Config) Copy() *Config {
	cp := *config
	return &cp
}
