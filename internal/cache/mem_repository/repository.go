package mem_repository

import (
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	defaultExpiration = time.Hour
	cleanupInterval   = time.Hour
)

type Repository struct {
	client *cache.Cache
}

func New() *Repository {
	return &Repository{
		client: cache.New(defaultExpiration, cleanupInterval),
	}
}
