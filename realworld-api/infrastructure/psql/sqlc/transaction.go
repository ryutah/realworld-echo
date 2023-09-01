package sqlc

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
)

type TransactionRunner struct {
	manager DBManager
}

func (t *TransactionRunner) Run(ctx context.Context, fn func(context.Context) error) (err error) {
	tx, err := t.manager.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if recov := recover(); recov != nil {
			err = t.rollback(ctx, tx, errors.Errorf("%v", recov))
		}
	}()

	ctx = t.manager.ContextWithExecutor(ctx, tx)
	if err := fn(ctx); err != nil {
		return t.rollback(ctx, tx, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return t.rollback(ctx, tx, errors.Wrap(err, "failed to commit transaction"))
	}
	return nil
}

func (t *TransactionRunner) rollback(ctx context.Context, tx pgx.Tx, err error) error {
	if tx == nil {
		return err
	}

	if rerr := tx.Rollback(ctx); rerr != nil {
		return errors.WithSecondaryError(err, rerr)
	}
	return err
}
