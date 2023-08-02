package cache

import (
	"github.com/lowl11/boost/internal/cache/mem_repository"
	"github.com/lowl11/boost/internal/cache/redis_repository"
	"github.com/lowl11/boost/pkg/enums/caches"
	"github.com/lowl11/boost/pkg/interfaces"
)

func getCacheRepository(cacheType string) (interfaces.CacheRepository, error) {
	switch cacheType {
	case caches.Memory:
		return mem_repository.New(), nil
	case caches.Redis:
		return redis_repository.New(), nil
	default:
		return nil, ErrorUndefinedCacheType(cacheType)
	}
}
