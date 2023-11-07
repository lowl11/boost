package elk_service

import (
	"github.com/lowl11/boost/data/enums/content_types"
	"github.com/lowl11/boost/pkg/web/requests"
	"time"
)

type Service struct {
	client *requests.Service
}

var instance *Service

func Get(host ...string) *Service {
	if instance != nil {
		return instance
	}

	var hostValue string
	if len(host) > 0 {
		hostValue = host[0]
	}

	instance = &Service{
		client: requests.
			New().
			SetBaseURL(hostValue).
			SetHeader("Content-Type", content_types.JSON).
			SetRetryCount(3).
			SetRetryWaitTime(time.Millisecond * 100),
	}
	return instance
}
