package mem_repository

import (
	"context"
	"github.com/lowl11/boost/pkg/types"
	"time"
)

func (repo Repository) All(_ context.Context) (map[string][]byte, error) {
	allCacheItems := repo.client.Items()
	all := make(map[string][]byte, len(allCacheItems))

	for key, value := range allCacheItems {
		all[key] = types.ToBytes(value)
	}

	return all, nil
}

func (repo Repository) Set(_ context.Context, key string, x any, expiration ...time.Duration) error {
	expires := defaultExpiration
	if len(expiration) > 0 {
		expires = expiration[0]
	}

	repo.client.Set(key, x, expires)
	return nil
}

func (repo Repository) Get(_ context.Context, key string) ([]byte, error) {
	x, found := repo.client.Get(key)
	if !found {
		return nil, nil
	}

	return types.ToBytes(x), nil
}

func (repo Repository) Delete(_ context.Context, key string) error {
	repo.client.Delete(key)
	return nil
}
