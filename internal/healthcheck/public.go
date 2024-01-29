package healthcheck

func (health *Healthcheck) Register(name, url string) {
	health.services = append(health.services, service{
		Name: name,
		URL:  url,
	})
}

func (health *Healthcheck) Trigger() error {
	for _, application := range health.services {
		if err := check(application); err != nil {
			return err
		}
	}

	return nil
}
