package redis_repository

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (repo Repository) allKeys(ctx context.Context) ([]string, error) {
	keys, err := repo.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, ErrorGetAllKeys(err)
	}

	return keys, nil
}

func (repo Repository) getType(ctx context.Context, key string) (string, error) {
	typeObject := repo.client.Type(ctx, key)
	if err := typeObject.Err(); err != nil {
		return "", ErrorRedisGetType(err)
	}

	return typeObject.Val(), nil
}

func (repo Repository) getList(ctx context.Context, key string) ([]byte, error) {
	result, err := repo.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, ErrorGetCacheByKey(key, err)
	}
	_ = result

	return nil, nil
}

func (repo Repository) getPrimitive(ctx context.Context, key string) ([]byte, error) {
	result, err := repo.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, ErrorGetCacheByKey(key, err)
	}

	return result, nil
}
