package request_helper

import "net/http"

func FillHeaders(request *http.Request, headers map[string]string) {
	if request == nil || headers == nil {
		return
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

func FillCookies(request *http.Request, cookies map[string]string) {
	if request == nil || cookies == nil {
		return
	}

	for name, value := range cookies {
		request.AddCookie(&http.Cookie{
			Name:  name,
			Value: value,
		})
	}
}
