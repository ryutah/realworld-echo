package usecase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/repository"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"go.uber.org/zap"
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

func (a *Article) Get(ctx context.Context, slugStr string) error {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()
	ctx = xlog.ContextWithLogFields(ctx, zap.String("slug", slugStr))

	xlog.Info(ctx, "Test")
	err := errors.New("test error")
	file, line, fn, _ := errors.GetOneLineSource(err)
	xerrorreport.NewErrorReporter("sample_service", "v1").Report(ctx, err, xerrorreport.ErrorContext{
		User: "sample_user",
		Location: xerrorreport.Location{
			File:     file,
			Line:     line,
			Function: fn,
		},
	})
	return a.outputPort.getArticle.OK(ctx, GetArticleResult{})
}
