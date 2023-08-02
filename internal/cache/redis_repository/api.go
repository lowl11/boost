package redis_repository

import "time"

func (repo Repository) All() map[string]any {
	panic("not implemented")
}

func (repo Repository) Set(key string, x any, expiration ...time.Duration) {
	panic("not implemented")
}

func (repo Repository) Get(key string) any {
	panic("not implemented")
}

func (repo Repository) Delete(key string) {
	panic("not implemented")
}
