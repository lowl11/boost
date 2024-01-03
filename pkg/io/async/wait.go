package async

import (
	"context"
)

func WaitAll(ctx context.Context, tasks ...func(ctx context.Context) error) []error {
	if len(tasks) == 0 {
		return nil
	}

	g := NewGroup(ctx)
	for _, f := range tasks {
		g.Run(f)
	}

	return g.
		Wait().
		Errors()
}
