package sql

import "context"

type Executor interface {
	Exec() error
}

type executor struct {
	ctx    context.Context
	query  string
	entity any
}

func newExecutor(ctx context.Context, query string, entity any) *executor {
	return &executor{
		ctx:    ctx,
		query:  query,
		entity: entity,
	}
}

func (e *executor) Exec() error {
	repo := getRepo()

	_, err := repo.DB(e.ctx).NamedExecContext(e.ctx, e.query, e.entity)
	if err != nil {
		return err
	}

	return nil
}
