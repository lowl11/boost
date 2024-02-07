package kafka

import (
	"context"
	"github.com/lowl11/boost/pkg/io/async"
	"sync"
	"time"
)

var (
	_snapshotLocks = sync.Map{}
)

// Snapshot - goes by every single record in topic which was created before Now.
// Function works for every topic in a thread safe way (i.e. "one topic" - "one mutex").
func Snapshot(ctx context.Context, cfg *Config, handler Handler, topic string, wait time.Duration) error {
	mx, ok := _snapshotLocks.Load(topic)
	if ok && mx != nil {
		mx.(*sync.Mutex).Lock()
		defer mx.(*sync.Mutex).Unlock()
	} else {
		_snapshotLocks.Store(topic, &sync.Mutex{})
	}

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
