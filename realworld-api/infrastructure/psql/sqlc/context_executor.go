package sqlc

import (
	"context"

	"github.com/cockroachdb/errors"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc/gen"
	"go.uber.org/fx"
)

type ContextExecutor interface {
	gen.DBTX
}

type DBManager interface {
	Querier(context.Context) gen.Querier
	ContextWithExecutor(context.Context, ContextExecutor) context.Context
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

type dbManager struct {
	dbpool *pgxpool.Pool
}

var _ DBManager = (*dbManager)(nil)

func NewDBManager(lc fx.Lifecycle) (*dbManager, error) {
	// TODO
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			dbpool.Close()
			return nil
		},
	})

	return &dbManager{
		dbpool: dbpool,
	}, nil
}

type contextExecutorKey struct{}

func (m *dbManager) Querier(ctx context.Context) gen.Querier {
	return gen.New(m.Executor(ctx))
}

func (m *dbManager) Executor(ctx context.Context) ContextExecutor {
	if exec, ok := ctx.Value(contextExecutorKey{}).(ContextExecutor); ok {
		return exec
	}
	return m.dbpool
}

func (m *dbManager) ContextWithExecutor(ctx context.Context, executor ContextExecutor) context.Context {
	return context.WithValue(ctx, contextExecutorKey{}, executor)
}

func (m *dbManager) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	tx, err := m.dbpool.BeginTx(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}
	return tx, nil
}
