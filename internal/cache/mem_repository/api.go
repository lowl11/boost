package mem_repository

import "time"

func (repo Repository) All() map[string]any {
	allCacheItems := repo.client.Items()
	all := make(map[string]any, len(allCacheItems))

	for key, value := range allCacheItems {
		all[key] = value
	}

	return all
}

func (repo Repository) Set(key string, x any, expiration ...time.Duration) {
	expires := defaultExpiration
	if len(expiration) > 0 {
		expires = expiration[0]
	}

	repo.client.Set(key, x, expires)
}

func (repo Repository) Get(key string) any {
	x, found := repo.client.Get(key)
	if !found {
		return nil
	}

	return x
}

func (repo Repository) Delete(key string) {
	repo.client.Delete(key)
}
