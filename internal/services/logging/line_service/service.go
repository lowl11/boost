package line_service

import (
	"sync"
)

type Service struct {
	logChannel       chan func()
	customLogChannel chan func()
	mutex            sync.Mutex
	customMutex      sync.Mutex
}

func New() *Service {
	line := &Service{
		logChannel:       make(chan func()),
		customLogChannel: make(chan func()),
	}

	go line.listen()
	go line.listenCustom()

	return line
}
