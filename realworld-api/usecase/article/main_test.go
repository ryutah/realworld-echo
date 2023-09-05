package article_test

import (
	"context"
	"testing"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mock_transaction "github.com/ryutah/realworld-echo/realworld-api/internal/mock/gen/domain/transaction"
	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/gen/usecase"
)

func mustNewTagName(s string) *model.TagName {
	tag, err := model.NewTagName(s)
	if err != nil {
		panic(err)
	}
	return &tag
}

type errorHandlerExpectationsOption[T any] struct {
	HandleArgsError     error
	HandleArgsOpts      []usecase.ErrorHandlerOption
	HandleReturnsResult *usecase.Result[T]
}

func transactionExpectations(t *testing.T, transaction *mock_transaction.MockTransaction) {
	t.Helper()

	transaction.EXPECT().
		Run(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, f func(context.Context) error) error {
			return f(ctx)
		})
}

func errorHandlerExpectations[T any](t *testing.T, errorHandler *mock_usecase.MockErrorHandler[T], opt errorHandlerExpectationsOption[T]) {
	t.Helper()

	errorHandler.EXPECT().
		Handle(
			mock.Anything,
			mock.Anything,
			lo.ToAnySlice(opt.HandleArgsOpts)...,
		).
		Run(
			func(ctx context.Context, err error, opts ...usecase.ErrorHandlerOption) {
				assert.ErrorIs(t, err, opt.HandleArgsError, "error of ErrorHandler#Handle args")
			},
		).
		Return(opt.HandleReturnsResult)
}
