package msgbus

import (
	"context"
	"github.com/lowl11/boost/data/enums/exchanges"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/cancel"
	"github.com/lowl11/boost/pkg/system/validator"
	"github.com/lowl11/boost/pkg/web/queue/rabbitmq/rmq"
	"github.com/lowl11/boost/pkg/web/queue/rabbitmq/rmq_connection"
)

const (
	defaultMessageBusExchangeName       = "message.bus"
	defaultMessageBusErrorsExchangeName = "message.bus.errors"
)

type DispatcherConfig struct {
	MessageBusExchangeName       string
	MessageBusErrorsExchangeName string
}

func defaultDispatcherConfig() DispatcherConfig {
	return DispatcherConfig{
		MessageBusExchangeName:       defaultMessageBusExchangeName,
		MessageBusErrorsExchangeName: defaultMessageBusErrorsExchangeName,
	}
}

type Dispatcher struct {
	validate       *validator.Validator
	skipValidation bool

	rmqService *rmq.Service

	messageBusExchangeName       string
	messageBusErrorsExchangeName string
}

func NewDispatcher(ctx context.Context, url string, cfg ...DispatcherConfig) (interfaces.Dispatcher, error) {
	validate, err := validator.New()
	if err != nil {
		return nil, err
	}

	rmq_connection.SetURL(url)
	_, err = rmq_connection.Get()
	if err != nil {
		return nil, err
	}

	var config DispatcherConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultDispatcherConfig()
	}

	dispatcher := &Dispatcher{
		validate:   validate,
		rmqService: rmq.New(),

		messageBusExchangeName:       config.MessageBusExchangeName,
		messageBusErrorsExchangeName: config.MessageBusErrorsExchangeName,
	}

	cancel.Get().Add()

	go func() {
		defer cancel.Get().Done()
		<-ctx.Done()
		if err = dispatcher.Close(); err != nil {
			log.Error("Close dispatcher error:", err)
			return
		}
		log.Info("Dispatcher closed")
	}()

	return dispatcher, nil
}

func (dispatcher *Dispatcher) Close() error {
	return dispatcher.rmqService.Close()
}

func (dispatcher *Dispatcher) Init() error {
	if err := dispatcher.declareExchanges(); err != nil {
		return err
	}

	return nil
}

func (dispatcher *Dispatcher) Dispatch(ctx context.Context, event any) error {
	eventName, err := nameOfEvent(event)
	if err != nil {
		return ErrorGetNameOfEvent(err)
	}

	if err = dispatcher.validateEvent(event); err != nil {
		return err
	}

	if err = dispatcher.rmqService.Publish(ctx, dispatcher.messageBusExchangeName, eventName, event); err != nil {
		return err
	}

	return nil
}

func (dispatcher *Dispatcher) SkipValidation() interfaces.Dispatcher {
	dispatcher.skipValidation = true
	return dispatcher
}

func (dispatcher *Dispatcher) validateEvent(event any) error {
	if dispatcher.skipValidation {
		return nil
	}

	return dispatcher.validate.Struct(event)
}

func (dispatcher *Dispatcher) declareExchanges() error {
	// message.bus
	if err := dispatcher.rmqService.NewExchange(dispatcher.messageBusExchangeName, exchanges.Direct); err != nil {
		return err
	}

	// message.bus.errors
	if err := dispatcher.rmqService.NewExchange(dispatcher.messageBusErrorsExchangeName, exchanges.Direct); err != nil {
		return err
	}

	return nil
}
