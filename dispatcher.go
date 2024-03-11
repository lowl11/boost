package boost

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/pkg/web/queue/msgbus"
)

// NewDispatcher creates new dispatcher instance for message bus
func NewDispatcher(amqpConnectionURL string) (Dispatcher, error) {
	dispatcher, err := msgbus.NewDispatcher(di.Get[App]().Context(), amqpConnectionURL)
	if err != nil {
		return nil, err
	}

	return dispatcher, nil
}

func RegisterDispatcher(connectionString string) {
	di.Register[Dispatcher](func() Dispatcher {
		dispatcher, err := NewDispatcher(connectionString)
		if err != nil {
			log.Fatal("Create dispatcher error:", err)
		}

		return dispatcher
	})
}
