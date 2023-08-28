package article

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	CreateArticleInputPort interface {
		Create(context.Context, *CreateArticleParam) *usecase.Result[CreateArticleResult]
	}
	CreateArticleParam  struct{}
	CreateArticleResult struct {
		Article model.Article
	}
)

type CreateArticle struct{}

func (c *CreateArticle) Create(ctx context.Context) {
}
