package cache

import (
	"github.com/lowl11/boost/internal/boosties/errors"
	"net/http"
)

const (
	typeErrorUndefinedCacheType = "UndefinedCacheType"
)

func ErrorUndefinedCacheType(cacheType string) error {
	return errors.
		New("Undefined cache type").
		SetType(typeErrorUndefinedCacheType).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("cache_type", cacheType)
}
