package elk_service

import (
	"github.com/lowl11/boost/data/enums/content_types"
	"github.com/lowl11/boost/pkg/web/requests"
	"time"
)

type Service struct {
	client *requests.Service
}

func New(host string) *Service {
	return &Service{
		client: requests.
			New().
			SetBaseURL(host).
			SetHeader("Content-Type", content_types.JSON).
			SetRetryCount(3).
			SetRetryWaitTime(time.Millisecond * 100),
	}
}
