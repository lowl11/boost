package sql

import (
	"context"
)

type Runner interface {
	Run(ctx context.Context, queries ...string) error
}

type runner struct {
	ctx     context.Context
	queries []string
}

func newRunner(ctx context.Context, queries []string) *runner {
	return &runner{
		ctx:     ctx,
		queries: queries,
	}
}

func Run(ctx context.Context, queries ...string) error {
	return newRunner(ctx, queries).Run()
}

func MustRun(ctx context.Context, queries ...string) {
	newRunner(ctx, queries).MustRun()
}

func RunError(ctx context.Context, handler func(err error), queries ...string) {
	newRunner(ctx, queries).RunError(handler)
}

func (r *runner) MustRun() {
	repo := getRepo()

	if r.queries == nil || len(r.queries) == 0 {
		return
	}

	for _, query := range r.queries {
		if _, err := repo.DB(r.ctx).ExecContext(r.ctx, query); err != nil {
			panic(err)
		}
	}
}

func (r *runner) Run() error {
	repo := getRepo()

	if r.queries == nil || len(r.queries) == 0 {
		return nil
	}

	for _, query := range r.queries {
		if _, err := repo.DB(r.ctx).ExecContext(r.ctx, query); err != nil {
			return err
		}
	}

	return nil
}

func (r *runner) RunError(handler func(err error)) {
	repo := getRepo()

	if r.queries == nil || len(r.queries) == 0 {
		return
	}

	for _, query := range r.queries {
		if _, err := repo.DB(r.ctx).ExecContext(r.ctx, query); err != nil {
			handler(err)
		}
	}
}
