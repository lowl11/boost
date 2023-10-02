package line_service

type InfoFunc func(args ...any)
type ErrorFunc func(err error, args ...any)

func (event *Service) AddInfo(logMessageFunc InfoFunc, args ...any) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	event.logChannel <- func() {
		logMessageFunc(args...)
	}
}

func (event *Service) AddInfoCustom(logMessageFunc InfoFunc, args ...any) {
	event.customMutex.Lock()
	defer event.customMutex.Unlock()

	event.customLogChannel <- func() {
		logMessageFunc(args...)
	}
}

func (event *Service) AddError(logMessageFunc ErrorFunc, err error, args ...any) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	event.logChannel <- func() {
		logMessageFunc(err, args...)
	}
}

func (event *Service) AddErrorCustom(logMessageFunc ErrorFunc, err error, args ...any) {
	event.customMutex.Lock()
	defer event.customMutex.Unlock()

	event.customLogChannel <- func() {
		logMessageFunc(err, args...)
	}
}
