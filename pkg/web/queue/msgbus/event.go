package msgbus

import (
	"github.com/lowl11/boost/pkg/system/types"
)

type Event struct {
	Name   string
	Action types.ListenerAction
	Object any
}
