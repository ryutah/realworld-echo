package article_test

import (
	"context"
	"testing"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/usecase"
)

func mustNewTag(s string) *model.ArticleTag {
	tag, err := model.NewArticleTag(s)
	if err != nil {
		panic(err)
	}
	return tag
}

type errorHandlerExpectationsOption[T any] struct {
	HandleArgsError      error
	HandleArgsOptsLength int
	HandleReturnsResult  *usecase.Result[T]
}

func errorHandlerExpectations[T any](t *testing.T, errorHandler *mock_usecase.MockErrorHandler[T], opt errorHandlerExpectationsOption[T]) {
	t.Helper()

	errorHandler.EXPECT().
		Handle(
			mock.Anything, mock.Anything, mock.Anything,
		).
		Run(
			func(ctx context.Context, err error, opts ...usecase.ErrorHandlerOption) {
				assert.ErrorIs(t, err, opt.HandleArgsError, "error of ErrorHandler#Handle args")
				assert.Len(t, opts, opt.HandleArgsOptsLength, "length of ErrorHandler#Hanel option args")
			},
		).
		Return(opt.HandleReturnsResult)
}
