package healthcheck

import (
	"net/http"
)

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
