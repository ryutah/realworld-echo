package article_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	mock_repository "github.com/ryutah/realworld-echo/realworld-api/internal/mock/repository"
	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetArticle_Get(t *testing.T) {
	type mocks_errorHandler struct {
		handle_args_error       error
		handle_args_opts_length int
		handle_returns_result   *usecase.Result[GetArticleResult]
	}
	type mocks_articleRepository struct {
		get_args_slugStr    model.Slug
		get_returns_article *model.Article
		get_returns_error   error
	}
	type mocks struct {
		errorHandler      mocks_errorHandler
		articleRepository mocks_articleRepository
	}
	type configs struct {
		errorHandler_handle_should_call bool
		article_get_should_skip         bool
	}
	type args struct {
		ctx     context.Context
		slugStr string
	}

	var (
		now        = premitive.NewJSTTime(time.Now())
		dummyError = errors.New("dummy")
		article1   = model.Article{
			Slug: "id",
			Contents: model.ArticleContents{
				Title:       "title",
				Description: "desc",
				Body:        "body",
			},
			Author:    "Author",
			CreatedAt: now,
			UpdatedAt: now,
		}
	)

	tests := []struct {
		name    string
		args    args
		mocks   mocks
		configs configs
		wants   *usecase.Result[GetArticleResult]
	}{
		{
			name: "when_given_any_slug_should_call_ArticleRepository_Get_and_GetArticleOutputPort_Success_and_return_nil",
			args: args{
				ctx:     context.TODO(),
				slugStr: "slug",
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					get_args_slugStr:    "slug",
					get_returns_article: &article1,
					get_returns_error:   nil,
				},
			},
			wants: usecase.Success(GetArticleResult{
				Article: article1,
			}),
		},
		{
			name: "when_given_not_exists_slug_should_call_ArticleRepositroy_Get_with_return_error_and_call_ErrorHandler_Handle_and_return_nil",
			args: args{
				ctx:     context.TODO(),
				slugStr: "notexists",
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 1,
					handle_returns_result:   usecase.Fail[GetArticleResult](usecase.NewFailResult(usecase.FailTypeNotFound, "error")),
				},
				articleRepository: mocks_articleRepository{
					get_args_slugStr:    "notexists",
					get_returns_article: nil,
					get_returns_error:   dummyError,
				},
			},
			wants: usecase.Fail[GetArticleResult](usecase.NewFailResult(usecase.FailTypeNotFound, "error")),
			configs: configs{
				errorHandler_handle_should_call: true,
			},
		},
		{
			name: "when_given_invalid_slug_should_call_ErrorHandler_and_return_nil",
			args: args{
				ctx:     context.TODO(),
				slugStr: strings.Repeat("a", 256),
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       derrors.Errors.Validation.Err,
					handle_args_opts_length: 1,
					handle_returns_result:   usecase.Fail[GetArticleResult](usecase.NewFailResult(usecase.FailTypeBadRequest, "error")),
				},
			},
			wants: usecase.Fail[GetArticleResult](usecase.NewFailResult(usecase.FailTypeBadRequest, "error")),
			configs: configs{
				errorHandler_handle_should_call: true,
				article_get_should_skip:         true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorHandler := mock_usecase.NewMockErrorHandler[GetArticleResult]()
			article := mock_repository.NewMockArticle()

			if tt.configs.errorHandler_handle_should_call {
				errorHandler.On(
					mock_usecase.ErrorHandlerFuncNames.Handle,
					mock.Anything, mock.Anything, mock.IsType([]usecase.ErrorHandlerOption{}),
				).Run(
					func(args mock.Arguments) {
						assert.ErrorIs(t, args.Error(1), tt.mocks.errorHandler.handle_args_error, "error of ErrorHandler#Handle args")
						if v, ok := args.Get(2).([]usecase.ErrorHandlerOption); ok {
							assert.Len(t, v, tt.mocks.errorHandler.handle_args_opts_length, "length of ErrorHandler#Hanel option args")
						}
					},
				).Return(
					tt.mocks.errorHandler.handle_returns_result,
				)
			}
			if !tt.configs.article_get_should_skip {
				article.On(
					mock_repository.ArticleFuncNames.Get,
					mock.Anything, tt.mocks.articleRepository.get_args_slugStr,
				).Return(
					tt.mocks.articleRepository.get_returns_article,
					tt.mocks.articleRepository.get_returns_error,
				)
			}

			a := NewGetArticle(errorHandler, article)
			result := a.Get(tt.args.ctx, tt.args.slugStr)

			assert.Equal(t, tt.wants, result)
			errorHandler.AssertExpectations(t)
			article.AssertExpectations(t)
		})
	}
}
