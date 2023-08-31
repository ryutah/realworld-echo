package sqlc

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

type TransactionRunner struct {
	manager DBManager
}

func (t *TransactionRunner) Run(ctx context.Context, fn func(context.Context) error) (err error) {
	tx, err := t.manager.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if recov := recover(); recov != nil {
			err = t.rollback(tx, errors.Errorf("%v", recov))
		}
	}()

	ctx = t.manager.ContextWithExecutor(ctx, tx)
	if err := fn(ctx); err != nil {
		return t.rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return t.rollback(tx, errors.Wrap(err, "failed to commit transaction"))
	}
	return nil
}

func (t *TransactionRunner) rollback(tx *sqlx.Tx, err error) error {
	if tx == nil {
		return err
	}

	if rerr := tx.Rollback(); rerr != nil {
		return errors.WithSecondaryError(err, rerr)
	}
	return err
}
