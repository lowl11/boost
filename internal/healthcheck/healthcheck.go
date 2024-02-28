package healthcheck

import (
	"github.com/lowl11/boost/pkg/io/exception"
	"net/http"
)

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

func (health *Healthcheck) Register(name, url string) {
	health.services = append(health.services, service{
		Name: name,
		URL:  url,
	})
}

func (health *Healthcheck) Trigger() error {
	for _, application := range health.services {
		if err := exception.Try(func() error {
			return check(application)
		}); err != nil {
			return err
		}
	}

	return nil
}

func check(application service) error {
	response, err := http.Get(application.URL)
	if err != nil {
		return errorHealthcheck(err)
	}

	if response.StatusCode != http.StatusOK {
		return errorHealthcheck(err)
	}

	return nil
}
