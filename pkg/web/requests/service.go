package requests

import (
	"crypto/tls"
	"github.com/lowl11/boost/log"
	"net/http"
	"net/url"
	"time"
)

type Service struct {
	baseURL string

	headers map[string]string
	cookies map[string]string

	timeout *time.Duration

	retryCount    int
	retryWaitTime time.Duration

	basicAuth *basicAuth

	client *http.Client
}

func New() *Service {
	return &Service{
		headers: make(map[string]string),
		cookies: make(map[string]string),

		retryWaitTime: time.Millisecond * 100,

		client: &http.Client{},
	}
}

func (service *Service) R() *Reqy {
	request := newReqy(service.baseURL, service.client).
		SetHeaders(service.headers).
		SetCookies(service.cookies).
		SetRetryCount(service.retryCount).
		SetRetryWaitTime(service.retryWaitTime)

	if service.timeout != nil {
		request.SetTimeout(*service.timeout)
	}

	if service.basicAuth != nil {
		request.SetBasicAuth(service.basicAuth)
	}

	return request
}

func (service *Service) SetBaseURL(baseURl string) *Service {
	service.baseURL = baseURl

	return service
}

func (service *Service) SetHeader(key, value string) *Service {
	service.headers[key] = value

	return service
}

func (service *Service) SetHeaders(headers map[string]string) *Service {
	for key, value := range headers {
		service.headers[key] = value
	}

	return service
}

func (service *Service) SetCookie(key, value string) *Service {
	service.cookies[key] = value

	return service
}

func (service *Service) SetCookies(cookies map[string]string) *Service {
	for key, value := range cookies {
		service.cookies[key] = value
	}

	return service
}

func (service *Service) SetTimeout(timeout time.Duration) *Service {
	service.timeout = &timeout
	return service
}

func (service *Service) SetTLSConfig(tls *tls.Config) *Service {
	transport := service.transport()
	transport.TLSClientConfig = tls
	service.client.Transport = transport

	return service
}

func (service *Service) SetProxy(proxyURL string) *Service {
	transport := service.transport()

	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		log.Error("Parse URL error:", err)
		return service
	}

	transport.Proxy = http.ProxyURL(parsedURL)
	service.client.Transport = transport

	return service
}

func (service *Service) SetRetryCount(count int) *Service {
	service.retryCount = count
	return service
}

func (service *Service) SetRetryWaitTime(waitTime time.Duration) *Service {
	service.retryWaitTime = waitTime
	return service
}

func (service *Service) SetBasicAuth(username, password string) *Service {
	service.basicAuth = &basicAuth{
		Username: username,
		Password: password,
	}
	return service
}

func (service *Service) transport() *http.Transport {
	transport, ok := service.client.Transport.(*http.Transport)
	if !ok {
		transport = &http.Transport{}
	}

	return transport
}
