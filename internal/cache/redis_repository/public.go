package redis_repository

import (
	"context"
	"github.com/lowl11/boost/pkg/enums/redis_types"
	"github.com/lowl11/boost/pkg/system/types"
	"time"
)

func (repo Repository) All(ctx context.Context) (map[string][]byte, error) {
	keys, err := repo.allKeys(ctx)
	if err != nil {
		return nil, err
	}

	all := make(map[string][]byte, len(keys))
	for _, key := range keys {
		value, err := repo.Get(ctx, key)
		if err != nil {
			return nil, err
		}

		all[key] = value
	}

	return all, nil
}

func (repo Repository) Search(ctx context.Context, pattern string) ([]string, error) {
	keys, err := repo.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, ErrorSearchKeys(pattern, err)
	}

	return keys, nil
}

func (repo Repository) Set(ctx context.Context, key string, x any, expiration ...time.Duration) error {
	expires := defaultExpiration
	if len(expiration) > 0 {
		expires = expiration[0]
	}

	if err := repo.client.Set(ctx, key, types.ToBytes(x), expires).Err(); err != nil {
		return ErrorSetCache(key, x, err)
	}

	return nil
}

func (repo Repository) Get(ctx context.Context, key string) ([]byte, error) {
	valueType, err := repo.getType(ctx, key)
	if err != nil {
		return nil, err
	}

	switch valueType {
	case redis_types.List:
		return repo.getList(ctx, key)
	case redis_types.String:
		return repo.getPrimitive(ctx, key)
	case redis_types.None:
		return nil, nil
	default:
		return nil, ErrorRedisUnknownType(key, valueType)
	}
}

func (repo Repository) Delete(ctx context.Context, key string) error {
	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return ErrorDeleteCache(key, err)
	}

	return nil
}

func (repo Repository) Close() error {
	return repo.client.Close()
}
