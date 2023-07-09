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
	GetResult struct {
		Article model.Article
	}
	GetInputPort[Ret any] interface {
		Get(ctx context.Context, slugStr string) (Ret, error)
	}
)

type Get[Ret any] struct {
	errorHandler usecase.GenericsErrorHandler[Ret]
	outputPort   usecase.NewOutputPort[GetResult, Ret]
	repository   struct {
		article repository.Article
	}
}

var _ GetInputPort[any] = (*Get[any])(nil)

func NewGet[Ret any](okPort usecase.NewOutputPort[GetResult, Ret], errorHandler usecase.GenericsErrorHandler[Ret], articleRepo repository.Article) *Get[Ret] {
	return &Get[Ret]{
		outputPort: okPort,
		repository: struct {
			article repository.Article
		}{
			article: articleRepo,
		},
	}
}

func (a *Get[Ret]) Get(ctx context.Context, slugStr string) (Ret, error) {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	ctx = xlog.ContextWithLogFields(ctx, zap.String("slug", slugStr))

	slug, err := model.NewSlug(slugStr)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithGenericsErrorRendrer(derrors.Errors.Validation.Err, usecase.GenericsBadRequest[Ret]))
	}

	xlog.Info(ctx, fmt.Sprintf("get article by: %v", slug))
	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithGenericsErrorRendrer(derrors.Errors.NotFound.Err, usecase.GenericsBadRequest[Ret]))
	}

	return a.outputPort.Success(ctx, GetResult{
		Article: *article,
	})
}

type ErrorType int

const (
	ErrorTypeInternalError ErrorType = iota + 1
	ErrorTypeNotFound
	ErrorTypeBadRequest
)

type Result[R any] struct {
	result      R
	errorResult *ErrorResult
}

func NewSuccessResult[R any](result R) *Result[R] {
	return &Result[R]{
		result: result,
	}
}

func NewFailResult[R any](result *ErrorResult) *Result[R] {
	return &Result[R]{
		errorResult: result,
	}
}

func (r *Result[R]) IsFailed() bool {
	return r.errorResult != nil
}

func (r *Result[R]) Success() R {
	return r.result
}

func (r *Result[R]) Fail() *ErrorResult {
	return r.errorResult
}

type GetResult2 struct {
	Article model.Article
}

type ErrorResult struct {
	Type         ErrorType
	Message      string
	Descriptions []string
}

func (a *Get[Ret]) Get2(ctx context.Context, slugStr string) *Result[GetResult2] {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	ctx = xlog.ContextWithLogFields(ctx, zap.String("slug", slugStr))

	slug, err := model.NewSlug(slugStr)
	if err != nil {
		return NewFailResult[GetResult2](&ErrorResult{})
	}

	xlog.Info(ctx, fmt.Sprintf("get article by: %v", slug))
	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return NewFailResult[GetResult2](&ErrorResult{})
	}

	return NewSuccessResult[GetResult2](GetResult2{
		Article: *article,
	})
}
