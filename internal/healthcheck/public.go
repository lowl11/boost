package healthcheck

import "github.com/lowl11/boost/pkg/io/exception"

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
