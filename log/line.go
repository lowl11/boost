package log

import "sync"

type line struct {
	logChannel       chan func()
	customLogChannel chan func()
	mutex            sync.Mutex
	customMutex      sync.Mutex
}

func newLine() *line {
	lineService := &line{
		logChannel:       make(chan func()),
		customLogChannel: make(chan func()),
	}

	go lineService.listen()
	go lineService.listenCustom()

	return lineService
}

func (service *line) AddInfo(logMessageFunc func(args ...any), args ...any) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.logChannel <- func() {
		logMessageFunc(args...)
	}
}

func (service *line) AddFormatInfo(logMessageFunc func(format string, args ...any), format string, args ...any) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.logChannel <- func() {
		logMessageFunc(format, args...)
	}
}

func (service *line) AddInfoCustom(logMessageFunc func(args ...any), args ...any) {
	service.customMutex.Lock()
	defer service.customMutex.Unlock()

	service.customLogChannel <- func() {
		logMessageFunc(args...)
	}
}

func (service *line) AddInfoFormatCustom(logMessageFunc func(format string, args ...any), format string, args ...any) {
	service.customMutex.Lock()
	defer service.customMutex.Unlock()

	service.customLogChannel <- func() {
		logMessageFunc(format, args...)
	}
}

func (service *line) AddError(logMessageFunc func(args ...any), args ...any) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.logChannel <- func() {
		logMessageFunc(args...)
	}
}

func (service *line) AddErrorCustom(logMessageFunc func(args ...any), args ...any) {
	service.customMutex.Lock()
	defer service.customMutex.Unlock()

	service.customLogChannel <- func() {
		logMessageFunc(args...)
	}
}

func (service *line) AddErrorFormatCustom(logMessageFunc func(format string, args ...any), format string, args ...any) {
	service.customMutex.Lock()
	defer service.customMutex.Unlock()

	service.customLogChannel <- func() {
		logMessageFunc(format, args...)
	}
}

func (service *line) listen() {
	for logMessageFunc := range service.logChannel {
		logMessageFunc()
	}
}

func (service *line) listenCustom() {
	for logMessageFunc := range service.customLogChannel {
		logMessageFunc()
	}
}
