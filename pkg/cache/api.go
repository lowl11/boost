package cache

import (
	"github.com/lowl11/boost/pkg/interfaces"
)

func New(cacheType string, cfg ...any) (interfaces.CacheRepository, error) {
	cacheRepo, err := getCacheRepository(cacheType, cfg...)
	if err != nil {
		return nil, err
	}

	return cacheRepo, nil
}
