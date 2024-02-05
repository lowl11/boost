package msgbus

import (
	"github.com/lowl11/boost/data/enums/exchanges"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/web/queue/rabbitmq/rmq"
	"github.com/lowl11/boost/pkg/web/queue/rabbitmq/rmq_connection"
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
	url        string
	rmqService *rmq.Service

	events []Event

	messageBusExchangeName       string
	messageBusErrorsExchangeName string
}

func NewListener(cfg ...ListenerConfig) interfaces.Listener {
	var config ListenerConfig
	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = defaultListenerConfig()
	}

	return &Listener{
		events: make([]Event, 0),

		messageBusExchangeName:       config.MessageBusExchangeName,
		messageBusErrorsExchangeName: config.MessageBusErrorsExchangeName,
	}
}

func (listener *Listener) Run(amqpConnectionURL string) error {
	listener.url = amqpConnectionURL

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
	return listener.rmqService.Close()
}

func (listener *Listener) Close() error {
	return listener.rmqService.Close()
}

func (listener *Listener) RegisterRoute(event Event) {
	listener.events = append(listener.events, event)
}

func (listener *Listener) Bind(event any, action func(ctx interfaces.EventContext) error) {
	eventName, err := nameOfEvent(event)
	if err != nil {
		return
	}

	listener.events = append(listener.events, Event{
		Name:   eventName,
		Action: action,
		Object: event,
	})
}

func (listener *Listener) EventsCount() int {
	return len(listener.events)
}

func (listener *Listener) declareExchanges() error {
	if err := listener.connect(); err != nil {
		return err
	}

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
	if err := listener.connect(); err != nil {
		return err
	}

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
		go listener.async(event, message)
	}
}

func (listener *Listener) async(event Event, message amqp.Delivery) {
	if err := event.Action(newContext(&message)); err != nil {
		log.Error("Event action error:", err)
		return
	}

	if err := listener.rmqService.Ack(message.DeliveryTag); err != nil {
		log.Error("Acknowledge message error:", err)
	}
}

func (listener *Listener) connect() error {
	if listener.rmqService != nil {
		return nil
	}

	rmq_connection.SetURL(listener.url)
	_, err := rmq_connection.Get()
	if err != nil {
		return err
	}

	listener.rmqService = rmq.New()
	return nil
}
