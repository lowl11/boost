package requests

import (
	"crypto/tls"
	"github.com/lowl11/boost/internal/services/reqy"
	"github.com/lowl11/lazylog/log"
	"net/http"
	"net/url"
	"time"
)

func (service *Service) R() *reqy.Request {
	request := reqy.
		NewRequest(service.baseURL, service.client).
		SetHeaders(service.headers).
		SetCookies(service.cookies)

	if service.timeout != nil {
		request.SetTimeout(*service.timeout)
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
	transport, err := service.transport()
	if err != nil {
		log.Error(err, errorGetHttpTransport)
		return service
	}

	transport.TLSClientConfig = tls

	return service
}

func (service *Service) SetProxy(proxyURL string) *Service {
	transport, err := service.transport()
	if err != nil {
		log.Error(err, errorGetHttpTransport)
		return service
	}

	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		log.Error(err, errorParseURL)
		return service
	}

	transport.Proxy = http.ProxyURL(parsedURL)

	return service
}
