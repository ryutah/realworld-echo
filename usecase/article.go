package usecase

import (
	"context"

	"github.com/ryutah/realworld-echo/domain/model"
	"github.com/ryutah/realworld-echo/domain/repository"
	"github.com/ryutah/realworld-echo/pkg/xlog"
)

type OKOutputPort[T any] interface {
	OK(context.Context, T) error
}

type GetArticleInputPort interface {
	Get(ctx context.Context, slugStr string) error
}

type GetArticleResult struct {
	Article *model.Article
}

type Article struct {
	outputPort struct {
		getArticle OKOutputPort[GetArticleResult]
		errors     ErrorOutputPort
	}
	repository struct {
		article repository.Article
	}
}

var _ GetArticleInputPort = (*Article)(nil)

func NewArticle(okPort OKOutputPort[GetArticleResult], errPort ErrorOutputPort) *Article {
	return &Article{
		outputPort: struct {
			getArticle OKOutputPort[GetArticleResult]
			errors     ErrorOutputPort
		}{
			getArticle: okPort,
			errors:     errPort,
		},
	}
}

func (a *Article) Get(ctx context.Context, slutStr string) error {
	xlog.Info(ctx, "Test!")
	return a.outputPort.getArticle.OK(ctx, GetArticleResult{})
}
