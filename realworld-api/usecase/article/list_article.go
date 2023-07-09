package article

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	ListArticleResult struct {
		Articles []model.Article
	}
	ListArticleInputPort interface {
		List(ctx context.Context) *usecase.Result[ListArticleResult]
	}
)

type ListArticle[Ret any] struct {
	errorHandler usecase.ErrorHandler[ListArticleResult]
	repository   struct {
		article repository.Article
	}
}

func (a *ListArticle[Ret]) List(ctx context.Context) *usecase.Result[ListArticleResult] {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	return usecase.Success(ListArticleResult{})
}
