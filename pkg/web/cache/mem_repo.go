package cache

import (
	"context"
	"github.com/lowl11/boost/pkg/io/types"
	"github.com/patrickmn/go-cache"
	"regexp"
	"strings"
	"time"
)

const (
	_memDefaultExpiration = time.Hour
	_memCleanupInterval   = time.Hour
)

type memRepo struct {
	client *cache.Cache
}

func newMemRepo() *memRepo {
	return &memRepo{
		client: cache.New(_memDefaultExpiration, _memCleanupInterval),
	}
}

func (repo memRepo) All(_ context.Context) (map[string][]byte, error) {
	allCacheItems := repo.client.Items()
	all := make(map[string][]byte, len(allCacheItems))

	for key, value := range allCacheItems {
		all[key] = types.ToBytes(value)
	}

	return all, nil
}

func (repo memRepo) Search(ctx context.Context, pattern string) ([]string, error) {
	all, err := repo.All(ctx)
	if err != nil {
		return nil, err
	}

	regexPattern := strings.ReplaceAll(pattern, "*", "(.*)")
	reg, _ := regexp.Compile(regexPattern)

	matchKeys := make([]string, 0)
	for key := range all {
		match := reg.FindAllString(key, -1)
		if len(match) > 0 && match[0] == key {
			matchKeys = append(matchKeys, key)
		}
	}

	return matchKeys, nil
}

func (repo memRepo) Refresh(_ context.Context, key string, expiration time.Duration) error {
	value, ok := repo.client.Get(key)
	if !ok {
		return nil
	}

	repo.client.Set(key, value, expiration)
	return nil
}

func (repo memRepo) Set(_ context.Context, key string, x any, expiration ...time.Duration) error {
	expires := _memDefaultExpiration
	if len(expiration) > 0 {
		expires = expiration[0]
	}

	repo.client.Set(key, x, expires)
	return nil
}

func (repo memRepo) Get(_ context.Context, key string) ([]byte, error) {
	x, found := repo.client.Get(key)
	if !found {
		return nil, nil
	}

	return types.ToBytes(x), nil
}

func (repo memRepo) Delete(_ context.Context, key string) error {
	repo.client.Delete(key)
	return nil
}

func (repo memRepo) Close() error {
	repo.client.Flush()
	return nil
}
