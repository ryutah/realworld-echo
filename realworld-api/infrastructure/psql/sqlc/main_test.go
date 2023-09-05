package sqlc_test

import (
	"context"
	"testing"

	mock_transaction "github.com/ryutah/realworld-echo/realworld-api/internal/mock/gen/domain/transaction"
	"github.com/stretchr/testify/mock"
)

func expectationsRunTransaction(t *testing.T, tx *mock_transaction.MockTransaction) {
	t.Helper()
	tx.EXPECT().
		Run(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, f func(context.Context) error) error {
			return f(ctx)
		})
}
