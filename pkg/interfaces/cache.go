package interfaces

import (
	"context"
	"time"
)

type CacheRepository interface {
	All(context.Context) (map[string][]byte, error)
	Set(context.Context, string, any, ...time.Duration) error
	Get(context.Context, string) ([]byte, error)
	Delete(context.Context, string) error

	Close() error
}
