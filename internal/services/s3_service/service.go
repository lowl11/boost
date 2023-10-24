package s3_service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	defaultMaxRetries = 3
)

type Service struct {
	maxRetries int
	region     string
	bucket     string

	connection *session.Session
	client     *s3.S3
}

var instance *Service

func Get() *Service {
	if instance != nil {
		return instance
	}

	instance = &Service{
		maxRetries: defaultMaxRetries,
	}
	return instance
}
