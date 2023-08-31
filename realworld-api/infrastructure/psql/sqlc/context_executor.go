package sqlc

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc/gen"
)

type ContextExecutor interface {
	gen.DBTX
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}

type DBManager interface {
	Querier(context.Context) gen.Querier
	ContextWithExecutor(context.Context, ContextExecutor) context.Context
	BeginTx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
}

type dbManager struct {
	db *sqlx.DB
}

var _ DBManager = (*dbManager)(nil)

func NewDBManager() *dbManager {
	// TODO
	return &dbManager{
		db: sqlx.MustOpen("postgresql", ""),
	}
}

type contextExecutorKey struct{}

func (m *dbManager) Querier(ctx context.Context) gen.Querier {
	return gen.New(m.executor(ctx))
}

func (m *dbManager) executor(ctx context.Context) ContextExecutor {
	if exec, ok := ctx.Value(contextExecutorKey{}).(ContextExecutor); ok {
		return exec
	}
	return m.db
}

func (m *dbManager) ContextWithExecutor(ctx context.Context, executor ContextExecutor) context.Context {
	return context.WithValue(ctx, contextExecutorKey{}, executor)
}

func (m *dbManager) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	tx, err := m.db.BeginTxx(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}
	return tx, nil
}
