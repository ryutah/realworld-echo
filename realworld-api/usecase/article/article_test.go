package article_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang/mock/gomock"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtesting"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article/mock"
)

func TestArticle_Get(t *testing.T) {
	type mocks_getArticleOutputPort struct {
		success_args_getArticleResult GetArticleResult
		success_returns_error         error
	}
	type mocks_errorHandler struct {
		handle_args_error       error
		handle_args_opts_length int
		handle_returns_error    error
	}
	type mocks_articleRepository struct {
		get_args_slugStr    model.Slug
		get_returns_article *model.Article
		get_returns_error   error
	}
	type mocks struct {
		getArticleOutputPort mocks_getArticleOutputPort
		errorHandler         mocks_errorHandler
		articleRepository    mocks_articleRepository
	}
	type configs struct {
		isError                 bool
		article_get_should_skip bool
	}
	type args struct {
		ctx     context.Context
		slugStr string
	}

	var (
		now        = time.Now()
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
		wants   error
	}{
		{
			name: "when_given_any_slug_should_call_ArticleRepository_Get_and_GetArticleOutputPort_Success_and_return_nil",
			args: args{
				ctx:     context.TODO(),
				slugStr: "slug",
			},
			mocks: mocks{
				getArticleOutputPort: mocks_getArticleOutputPort{
					success_args_getArticleResult: GetArticleResult{
						Article: article1,
					},
					success_returns_error: nil,
				},
				articleRepository: mocks_articleRepository{
					get_args_slugStr:    "slug",
					get_returns_article: &article1,
					get_returns_error:   nil,
				},
			},
			wants: nil,
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
					handle_returns_error:    nil,
				},
				articleRepository: mocks_articleRepository{
					get_args_slugStr:    "notexists",
					get_returns_article: nil,
					get_returns_error:   dummyError,
				},
			},
			configs: configs{
				isError: true,
			},
			wants: nil,
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
					handle_returns_error:    nil,
				},
			},
			configs: configs{
				isError:                 true,
				article_get_should_skip: true,
			},
			wants: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			getOutputPort := mock.NewMockGetArticleOutputPort(ctrl)
			errorHandler := mock.NewMockErrorHandler(ctrl)
			article := mock.NewMockArticle(ctrl)

			if tt.configs.isError {
				errorHandler.EXPECT().
					Handle(gomock.Any(), xtesting.MatchError(tt.mocks.errorHandler.handle_args_error), gomock.Len(tt.mocks.errorHandler.handle_args_opts_length)).
					Return(tt.mocks.errorHandler.handle_returns_error)
			} else {
				getOutputPort.EXPECT().
					Success(gomock.Any(), tt.mocks.getArticleOutputPort.success_args_getArticleResult).
					Return(tt.mocks.getArticleOutputPort.success_returns_error)
			}
			if !tt.configs.article_get_should_skip {
				article.EXPECT().
					Get(gomock.Any(), tt.mocks.articleRepository.get_args_slugStr).
					Return(tt.mocks.articleRepository.get_returns_article, tt.mocks.articleRepository.get_returns_error)
			}

			a := NewArticle(getOutputPort, errorHandler, article)
			err := a.Get(tt.args.ctx, tt.args.slugStr)

			xtesting.CompareError(t, "Get", tt.wants, err)
		})
	}
}
