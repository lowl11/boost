package healthcheck

type Healthcheck struct {
	services []service
}

type service struct {
	Name string
	URL  string
}

func New() *Healthcheck {
	return &Healthcheck{
		services: make([]service, 0),
	}
}
