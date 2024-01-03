package boost

import (
	"github.com/lowl11/boost/pkg/web/queue/msgbus"
)

// NewDispatcher creates new dispatcher instance for message bus
func NewDispatcher(amqpConnectionURL string) (Dispatcher, error) {
	dispatcher, err := msgbus.NewDispatcher(amqpConnectionURL)
	if err != nil {
		return nil, err
	}

	return dispatcher, nil
}
