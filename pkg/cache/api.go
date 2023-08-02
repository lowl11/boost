package cache

import "github.com/lowl11/boost/pkg/interfaces"

func New(cacheType string) (interfaces.CacheRepository, error) {
	cacheRepo, err := getCacheRepository(cacheType)
	if err != nil {
		return nil, err
	}

	return cacheRepo, nil
}
