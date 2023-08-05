package msgbus

import (
	"github.com/lowl11/boost/internal/helpers/event_helper"
	"github.com/lowl11/boost/internal/queue/event_context"
	"github.com/lowl11/boost/internal/queue/rmq_service"
	"github.com/lowl11/boost/pkg/enums/exchanges"
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/lazylog/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ListenerConfig struct {
	MessageBusExchangeName       string
	MessageBusErrorsExchangeName string
}

func defaultListenerConfig() ListenerConfig {
	return ListenerConfig{
		MessageBusExchangeName:       defaultMessageBusExchangeName,
		MessageBusErrorsExchangeName: defaultMessageBusErrorsExchangeName,
	}
}

type Listener struct {
	rmqService *rmq_service.Service

	events []Event

	messageBusExchangeName       string
	messageBusErrorsExchangeName string
}

func NewListener(url string, cfg ...ListenerConfig) (interfaces.Listener, error) {
	var config ListenerConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultListenerConfig()
	}

	rmqService, err := rmq_service.New(url)
	if err != nil {
		return nil, err
	}

	return &Listener{
		rmqService: rmqService,
		events:     make([]Event, 0),

		messageBusExchangeName:       config.MessageBusExchangeName,
		messageBusErrorsExchangeName: config.MessageBusErrorsExchangeName,
	}, nil
}

func (listener *Listener) Run() error {
	if err := listener.declareExchanges(); err != nil {
		return err
	}

	if err := listener.declareEvents(); err != nil {
		return err
	}

	for _, event := range listener.events {
		messages, err := listener.rmqService.Consume(event.Name)
		if err != nil {
			return err
		}

		go listener.listen(messages, event)
	}

	infinite := make(chan struct{})
	<-infinite
	return nil
}

func (listener *Listener) Bind(event any, action func(ctx interfaces.EventContext) error) {
	eventName, err := event_helper.NameOfEvent(event)
	if err != nil {
		return
	}

	listener.events = append(listener.events, Event{
		Name:   eventName,
		Action: action,
		Object: event,
	})
}

func (listener *Listener) declareExchanges() error {
	// message.bus
	if err := listener.rmqService.NewExchange(listener.messageBusExchangeName, exchanges.Direct); err != nil {
		return err
	}

	// message.bus.errors
	if err := listener.rmqService.NewExchange(listener.messageBusErrorsExchangeName, exchanges.Direct); err != nil {
		return err
	}

	return nil
}

func (listener *Listener) declareEvents() error {
	for _, event := range listener.events {
		// first, create queue
		if _, err := listener.rmqService.NewQueue(event.Name); err != nil {
			return err
		}

		// after this, bind exchange & queue
		if err := listener.rmqService.Bind(listener.messageBusExchangeName, event.Name); err != nil {
			return err
		}
	}

	return nil
}

func (listener *Listener) listen(messages <-chan amqp.Delivery, event Event) {
	for message := range messages {
		if err := event.Action(event_context.New(&message)); err != nil {
			log.Error(err, "Event action error")
			continue
		}

		if err := listener.rmqService.Ack(message.DeliveryTag); err != nil {
			log.Error(err, "Acknowledge message error")
		}
	}
}
