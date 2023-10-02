package types

import (
	"github.com/lowl11/boost/data/interfaces"
)

type ListenerAction func(event interfaces.EventContext) error
