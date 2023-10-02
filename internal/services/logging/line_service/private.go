package line_service

func (event *Service) listen() {
	for logMessageFunc := range event.logChannel {
		logMessageFunc()
	}
}

func (event *Service) listenCustom() {
	for logMessageFunc := range event.customLogChannel {
		logMessageFunc()
	}
}
