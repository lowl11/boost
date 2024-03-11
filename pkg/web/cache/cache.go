package cache

import (
	"context"
	"github.com/lowl11/boost/data/enums/caches"
	"github.com/lowl11/boost/data/interfaces"
)

type RedisConfig struct {
	URL      string
	Password string
	DB       int
}

func New(ctx context.Context, cacheType string, cfg ...any) (interfaces.Cache, error) {
	cacheRepo, err := getCacheRepository(ctx, cacheType, cfg...)
	if err != nil {
		return nil, err
	}

	return cacheRepo, nil
}

func getCacheRepository(ctx context.Context, cacheType string, cfg ...any) (interfaces.Cache, error) {
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

		return newRedisRepo(ctx, redisConfigInstance{
			URL:      redisConfig.URL,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
	default:
		return nil, ErrorUndefinedCacheType(cacheType)
	}
}
