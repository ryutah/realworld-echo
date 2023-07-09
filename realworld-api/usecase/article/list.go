package article

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	ListResult struct {
		Articles []model.Article
	}
	ListInputPort[Ret any] interface {
		List(ctx context.Context) (Ret, error)
	}
)

type List[Ret any] struct {
	repository struct {
		article repository.Article
	}
	outputPort usecase.NewOutputPort[ListResult, Ret]
}

func (a *List[Ret]) List(ctx context.Context) (Ret, error) {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()
	a.repository.article.Get(ctx, "")

	return a.outputPort.Success(ctx, ListResult{
		Articles: nil,
	})
}
