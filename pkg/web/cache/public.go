package cache

import (
	"github.com/lowl11/boost/data/interfaces"
)

type RedisConfig struct {
	URL      string
	Password string
	DB       int
}

func New(cacheType string, cfg ...any) (interfaces.CacheRepository, error) {
	cacheRepo, err := getCacheRepository(cacheType, cfg...)
	if err != nil {
		return nil, err
	}

	return cacheRepo, nil
}
