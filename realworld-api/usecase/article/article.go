package article

import (
	"context"
	"fmt"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"go.uber.org/zap"
)

type (
	GetArticleResult struct {
		Article model.Article
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
		Article model.Article
	}
	CreateArticleInputPort interface {
		Create(context.Context, CreateArticleRequest) error
	}
)

// NOTE: workfaround for gomock
//
//	see: https://github.com/golang/mock/issues/621#issuecomment-1094351718
type (
	GetArticleOutputPort    = usecase.OutputPort[GetArticleResult]
	CreateArticleOutputPort = usecase.OutputPort[CreateArticleResult]
)

type Article struct {
	errorHandler usecase.ErrorHandler
	outputPort   struct {
		get    GetArticleOutputPort
		create CreateArticleOutputPort
	}
	repository struct {
		article repository.Article
	}
}

var _ GetArticleInputPort = (*Article)(nil)

func NewArticle(okPort GetArticleOutputPort, errorHandler usecase.ErrorHandler, articleRepo repository.Article) *Article {
	return &Article{
		errorHandler: errorHandler,
		outputPort: struct {
			get    GetArticleOutputPort
			create CreateArticleOutputPort
		}{
			get: okPort,
		},
		repository: struct {
			article repository.Article
		}{
			article: articleRepo,
		},
	}
}

func (a *Article) Get(ctx context.Context, slugStr string) error {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	ctx = xlog.ContextWithLogFields(ctx, zap.String("slug", slugStr))

	slug, err := model.NewSlug(slugStr)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithErrorRendrer(derrors.Errors.Validation.Err, usecase.BadRequest))
	}

	xlog.Info(ctx, fmt.Sprintf("get article by: %v", slug))
	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithErrorRendrer(derrors.Errors.NotFound.Err, usecase.NotFound))
	}

	return a.outputPort.get.Success(ctx, GetArticleResult{
		Article: *article,
	})
}

func (a *Article) Create(ctx context.Context, req CreateArticleRequest) error {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	return nil
}
