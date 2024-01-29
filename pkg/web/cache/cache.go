package cache

import (
	"github.com/lowl11/boost/data/enums/caches"
	"github.com/lowl11/boost/data/interfaces"
)

func getCacheRepository(cacheType string, cfg ...any) (interfaces.CacheRepository, error) {
	switch cacheType {
	case caches.Memory:
		return newMemRepo(), nil
	case caches.Redis:
		var redisConfig RedisConfig
		if len(cfg) > 0 {
			rdsCfg, ok := cfg[0].(RedisConfig)
			if !ok {
				return nil, ErrorRedisConfigRequired()
			}

			redisConfig = rdsCfg
		}

		return newRedisRepo(redisConfigInstance{
			URL:      redisConfig.URL,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
	default:
		return nil, ErrorUndefinedCacheType(cacheType)
	}
}
