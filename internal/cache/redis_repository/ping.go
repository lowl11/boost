package redis_repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func ping(client *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return ErrorRedisPing(err)
	}

	return nil
}
