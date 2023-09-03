package sqlc

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
)

type RawSelector[T any] interface {
	Select(context.Context, ContextExecutor, squirrel.SelectBuilder) ([]T, error)
}

type rawSelector[T any] struct{}

func NewRawSelector[T any]() RawSelector[T] {
	return &rawSelector[T]{}
}

func (s *rawSelector[T]) Select(ctx context.Context, exec ContextExecutor, builder squirrel.SelectBuilder) ([]T, error) {
	q, param, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}
	rows, err := exec.Query(ctx, q, param...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exec query")
	}
	result, err := pgx.CollectRows[T](rows, pgx.RowToStructByName[T])
	if err != nil {
		return nil, errors.Wrap(err, "failed to collect rows")
	}
	return result, nil
}
