package cache

import (
	"github.com/lowl11/boost/internal/cache/mem_repository"
	"github.com/lowl11/boost/internal/cache/redis_repository"
	"github.com/lowl11/boost/pkg/enums/caches"
	"github.com/lowl11/boost/pkg/interfaces"
)

func getCacheRepository(cacheType string, cfg ...any) (interfaces.CacheRepository, error) {
	switch cacheType {
	case caches.Memory:
		return mem_repository.New(), nil
	case caches.Redis:
		var redisConfig RedisConfig
		if len(cfg) > 0 {
			rdsCfg, ok := cfg[0].(RedisConfig)
			if !ok {
				return nil, ErrorRedisConfigRequired()
			}

			redisConfig = rdsCfg
		}

		return redis_repository.New(redis_repository.Config{
			URL:      redisConfig.URL,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
	default:
		return nil, ErrorUndefinedCacheType(cacheType)
	}
}
