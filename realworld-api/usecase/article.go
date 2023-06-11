package usecase

import (
	"context"

	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/domain/repository"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"go.uber.org/zap"
)

type (
	GetArticleResult struct {
		Article *model.Article
	}
	GetArticleInputPort interface {
		Get(ctx context.Context, slugStr string) error
	}
)

type (
	CreateArticleRequest struct {
		Title       string
		Description string
		Body        string
		TagList     []string
	}
	CreateArticleResult struct {
		Article *model.Article
	}
	CreateArticleInputPort interface {
		Create(context.Context, CreateArticleRequest) error
	}
)

// NOTE: workfaround for gomock
//
//	see: https://github.com/golang/mock/issues/621#issuecomment-1094351718
type (
	GetArticleOutputPort    = OutputPort[GetArticleResult]
	CreateArticleOutputPort = OutputPort[CreateArticleResult]
)

type Article struct {
	errorHandler ErrorHandler
	outputPort   struct {
		get    GetArticleOutputPort
		create CreateArticleOutputPort
	}
	repository struct {
		article repository.Article
	}
}

var _ GetArticleInputPort = (*Article)(nil)

func NewArticle(okPort GetArticleOutputPort, errorHandler ErrorHandler) *Article {
	return &Article{
		errorHandler: errorHandler,
		outputPort: struct {
			get    GetArticleOutputPort
			create CreateArticleOutputPort
		}{
			get: okPort,
		},
	}
}

func (a *Article) Get(ctx context.Context, slugStr string) error {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	ctx = xlog.ContextWithLogFields(ctx, zap.String("slug", slugStr))

	slug, err := premitive.NewSlug(slugStr)
	if err != nil {
		return a.errorHandler.handle(ctx, err, withErrorRendrer(derrors.Errors.Validation.Err, badRequest))
	}

	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return a.errorHandler.handle(ctx, err, withErrorRendrer(derrors.Errors.NotFound.Err, notFound))
	}

	return a.outputPort.get.Success(ctx, GetArticleResult{
		Article: article,
	})
}

func (a *Article) Create(ctx context.Context, req CreateArticleRequest) error {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	return nil
}
