package interfaces

import (
	"context"
	"time"
)

type Cache interface {
	All(context.Context) (map[string][]byte, error)
	Search(context.Context, string) ([]string, error)
	Refresh(ctx context.Context, key string, expiration time.Duration) error
	Set(context.Context, string, any, ...time.Duration) error
	Get(context.Context, string) ([]byte, error)
	Delete(context.Context, string) error

	Close() error
}
