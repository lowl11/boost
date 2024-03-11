package cache

import (
	"context"
	"github.com/lowl11/boost/data/enums/redis_types"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/cancel"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	_redisDefaultExpiration = time.Hour
	_redisCleanupInterval   = time.Hour
)

type redisConfigInstance struct {
	URL      string
	Password string
	DB       int
}

type redisRepo struct {
	client *redis.Client
}

func newRedisRepo(ctx context.Context, cfg redisConfigInstance) (*redisRepo, error) {
	if cfg.URL == "" {
		return nil, errors.New("URL is required").
			SetType("Redis_URLRequired")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.URL,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := pingRedis(client); err != nil {
		return nil, err
	}

	cancel.Get().Add()

	go func() {
		defer cancel.Get().Done()
		<-ctx.Done()
		if err := client.Close(); err != nil {
			log.Error("Close redis connection error:", err)
			return
		}
		log.Info("Redis closed")
	}()

	return &redisRepo{
		client: client,
	}, nil
}

func (repo redisRepo) All(ctx context.Context) (map[string][]byte, error) {
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

func (repo redisRepo) Search(ctx context.Context, pattern string) ([]string, error) {
	keys, err := repo.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, errors.New("Search keys by pattern error").
			SetType("Redis_SearchKeys").
			SetError(err).
			AddContext("pattern", pattern)
	}

	return keys, nil
}

func (repo redisRepo) Refresh(ctx context.Context, key string, expiration time.Duration) error {
	if err := repo.client.Expire(ctx, key, expiration).Err(); err != nil {
		return errors.
			New("Refresh expiration of key").
			SetType("Redis_RefreshExpire").
			SetError(err)
	}

	return nil
}

func (repo redisRepo) Set(ctx context.Context, key string, x any, expiration ...time.Duration) error {
	expires := _redisDefaultExpiration
	if len(expiration) > 0 {
		expires = expiration[0]
	}

	if err := repo.client.Set(ctx, key, types.ToBytes(x), expires).Err(); err != nil {
		return errors.New("Set cache error").
			SetType("Redis_SetCache").
			SetError(err).
			SetContext(map[string]any{
				"key":   key,
				"value": x,
			})
	}

	return nil
}

func (repo redisRepo) Get(ctx context.Context, key string) ([]byte, error) {
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
		return nil, errors.New("Unknown Redis type").
			SetType("Redis_UnknownType").
			AddContext("redis_type", valueType).
			AddContext("key", key)
	}
}

func (repo redisRepo) Delete(ctx context.Context, key string) error {
	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return errors.New("Delete cache error").
			SetType("Redis_DeleteCache").
			SetError(err).
			AddContext("key", key)
	}

	return nil
}

func (repo redisRepo) Close() error {
	return repo.client.Close()
}

func pingRedis(client *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return errors.New("Ping Redis error").
			SetType("Redis_Ping").
			SetError(err)
	}

	return nil
}

func (repo redisRepo) allKeys(ctx context.Context) ([]string, error) {
	keys, err := repo.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, errors.New("Get all keys error").
			SetType("Redis_GetAllKeys").
			SetError(err)
	}

	return keys, nil
}

func (repo redisRepo) getType(ctx context.Context, key string) (string, error) {
	typeObject := repo.client.Type(ctx, key)
	if err := typeObject.Err(); err != nil {
		return "", errors.New("Get key type error").
			SetType("Redis_GetType").
			SetError(err)
	}

	return typeObject.Val(), nil
}

func (repo redisRepo) getList(ctx context.Context, key string) ([]byte, error) {
	result, err := repo.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, errors.New("Get cache by key error").
			SetType("Redis_GetCacheByKey").
			SetError(err).
			AddContext("key", key)
	}
	_ = result

	return nil, nil
}

func (repo redisRepo) getPrimitive(ctx context.Context, key string) ([]byte, error) {
	result, err := repo.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, errors.New("Get cache by key error").
			SetType("Redis_GetCacheByKey").
			SetError(err).
			AddContext("key", key)
	}

	return result, nil
}
