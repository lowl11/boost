package types

import "github.com/lowl11/boost/pkg/interfaces"

type ListenerAction func(event interfaces.EventContext) error
