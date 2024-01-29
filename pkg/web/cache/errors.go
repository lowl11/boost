package cache

import (
	"github.com/lowl11/boost/errors"
	"net/http"
)

const (
	typeErrorUndefinedCacheType  = "UndefinedCacheType"
	typeErrorRedisConfigRequired = "RedisConfigRequired"
)

func ErrorUndefinedCacheType(cacheType string) error {
	return errors.
		New("Undefined cache type").
		SetType(typeErrorUndefinedCacheType).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("cache_type", cacheType)
}

func ErrorRedisConfigRequired() error {
	return errors.
		New("Redis config is required").
		SetType(typeErrorRedisConfigRequired).
		SetHttpCode(http.StatusInternalServerError)
}
