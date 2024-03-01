package boost

import (
	"github.com/valyala/fasthttp"
	"time"
)

func CookieOption(expiresAt time.Time, httpOnly bool, domain string) func(cookie *fasthttp.Cookie) {
	return func(cookie *fasthttp.Cookie) {
		CookieWithExpireAt(expiresAt)
		CookieWithHttpOnly(httpOnly)
		CookieWithDomain(domain)
	}
}

func CookieWithHttpOnly(httpOnly bool) func(cookie *fasthttp.Cookie) {
	return func(cookie *fasthttp.Cookie) {
		cookie.SetHTTPOnly(httpOnly)
	}
}

func CookieWithDomain(domain string) func(cookie *fasthttp.Cookie) {
	return func(cookie *fasthttp.Cookie) {
		cookie.SetDomain(domain)
	}
}

func CookieWithExpireAt(expireAt time.Time) func(cookie *fasthttp.Cookie) {
	return func(cookie *fasthttp.Cookie) {
		cookie.SetExpire(expireAt)
	}
}
