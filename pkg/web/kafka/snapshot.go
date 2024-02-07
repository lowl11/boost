package kafka

import (
	"context"
	"github.com/lowl11/boost/pkg/io/async"
	"time"
)

func Snapshot(ctx context.Context, cfg *Config, handler Handler, topic string, wait time.Duration) error {
	cfg = cfg.
		Copy().
		WithAutocommit(time.Second).
		WithOffset(Oldest)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c, err := NewConsumer(ctx, cfg)
	if err != nil {
		return err
	}

	task := async.NewTask(context.Background()).Run(func(ctx context.Context) error {
		time.Sleep(wait)
		cancel()
		return nil
	})

	if err = c.StartListening(topic, handler); err != nil {
		return err
	}

	return task.Wait()
}
