package usecase

import (
	"context"

	"github.com/ryutah/realworld-echo/domain/model"
	"github.com/ryutah/realworld-echo/domain/repository"
	"github.com/ryutah/realworld-echo/internal/xlog"
)

type ErrorResult struct {
	Message      string
	Descriptions []string
}

type ErrorOutputPort interface {
	InternalError(context.Context, ErrorResult) error
	NotFound(context.Context, ErrorResult) error
}

type GetArticleInputPort interface {
	Get(ctx context.Context, slugStr string) error
}

type GetArticleOutputPort interface {
	ErrorOutputPort
	Ok(context.Context, GetArticleResult) error
}

type GetArticleResult struct {
	Article *model.Article
}

type Article struct {
	outputPort struct {
		getArticle GetArticleOutputPort
	}
	repository struct {
		article repository.Article
	}
}

var _ GetArticleInputPort = (*Article)(nil)

func NewArticle(getArticle GetArticleOutputPort) *Article {
	return &Article{
		outputPort: struct {
			getArticle GetArticleOutputPort
		}{
			getArticle: getArticle,
		},
	}
}

func (a *Article) Get(ctx context.Context, slutStr string) error {
	xlog.Info(ctx, "Test!")
	return a.outputPort.getArticle.Ok(ctx, GetArticleResult{})
}
