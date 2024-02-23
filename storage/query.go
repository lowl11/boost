package storage

import (
	"context"
)

var (
	_logEnabled = false
)

func EnableLog() {
	_logEnabled = true
}

type Query interface {
	//
}

type queryable struct {
	//
}

func (q *queryable) All(ctx context.Context) error {
	return nil
}
