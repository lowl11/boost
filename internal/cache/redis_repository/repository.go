package redis_repository

import (
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	defaultExpiration = time.Hour
	cleanupInterval   = time.Hour
)

type Config struct {
	URL      string
	Password string
	DB       int
}

type Repository struct {
	client *redis.Client
}

func New(cfg Config) (*Repository, error) {
	if cfg.URL == "" {
		return nil, ErrorURLRequired()
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.URL,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := ping(client); err != nil {
		return nil, err
	}

	return &Repository{
		client: client,
	}, nil
}
