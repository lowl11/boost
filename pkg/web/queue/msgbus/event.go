package msgbus

import (
	"github.com/lowl11/boost/data/domain"
)

type Event struct {
	Name   string
	Action domain.ListenerAction
	Object any
}
