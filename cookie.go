package boost

import (
	"github.com/valyala/fasthttp"
	"time"
)

func CookieOption(expiresAt time.Time, httpOnly bool, path string) func(cookie *fasthttp.Cookie) {
	return func(cookie *fasthttp.Cookie) {
		CookieWithExpireAt(expiresAt)
		CookieWithHttpOnly(httpOnly)
		CookieWithPath(path)
	}
}

func CookieWithPath(path string) func(cookie *fasthttp.Cookie) {
	return func(cookie *fasthttp.Cookie) {
		cookie.SetPath(path)
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
