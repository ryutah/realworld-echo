package article_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	mock_repository "github.com/ryutah/realworld-echo/realworld-api/internal/mock/article/repository"
	mock_auth_service "github.com/ryutah/realworld-echo/realworld-api/internal/mock/auth/service"
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
	type mocks_favoriteRepository struct {
		listBySlug_args_slug         model.Slug
		listBySlug_returns_favorites model.FavoriteSlice
		listBySlug_returns_error     error
	}
	type mocks_authService struct {
		currentUser_returns_user  *authmodel.User
		currentUser_returns_error error
	}
	type mocks struct {
		errorHandler       mocks_errorHandler
		articleRepository  mocks_articleRepository
		favoriteRepository mocks_favoriteRepository
		authService        mocks_authService
	}
	type configs struct {
		errorHandler_handle_should_call bool
		article_get_should_skip         bool
		favorite_listBySlug_should_skip bool
		auth_currentUser_should_skip    bool
	}
	type args struct {
		slugStr string
	}

	var (
		now                         = premitive.NewJSTTime(time.Now())
		dummyError                  = errors.New("dummy")
		slug1      model.Slug       = "slug"
		user1      authmodel.UserID = "user1"
		user2      authmodel.UserID = "user2"
		failResult                  = usecase.Fail[GetArticleResult](usecase.NewFailResult(usecase.FailTypeInternalError, "error"))
		testData1                   = struct {
			args  args
			mocks mocks
			wants *usecase.Result[GetArticleResult]
		}{
			args: args{
				slugStr: slug1.String(),
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					get_args_slugStr: slug1,
					get_returns_article: &model.Article{
						Slug: slug1,
						Contents: model.ArticleContents{
							Title:       "title",
							Description: "desc",
							Body:        "body",
						},
						Author:    "Author",
						CreatedAt: now,
						UpdatedAt: now,
					},
					get_returns_error: nil,
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlug_args_slug: slug1,
					listBySlug_returns_favorites: model.FavoriteSlice{
						{ArticleSlug: slug1, UserID: user1},
						{ArticleSlug: slug1, UserID: user2},
					},
				},
				authService: mocks_authService{
					currentUser_returns_user: &authmodel.User{ID: user1},
				},
			},
			wants: usecase.Success(GetArticleResult{
				Article: model.Article{
					Slug: slug1,
					Contents: model.ArticleContents{
						Title:       "title",
						Description: "desc",
						Body:        "body",
					},
					Author:    "Author",
					CreatedAt: now,
					UpdatedAt: now,
				},
				Favorited: true,
				Favorites: model.FavoriteSlice{
					{ArticleSlug: slug1, UserID: user1},
					{ArticleSlug: slug1, UserID: user2},
				},
			}),
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
			name:  "when_given_any_slug_should_call_ArticleRepository_Get_and_GetArticleOutputPort_Success_and_return_nil",
			args:  testData1.args,
			mocks: testData1.mocks,
			wants: testData1.wants,
		},
		{
			name: "when_given_any_slug_without_signedin_should_call_ArticleRepository_Get_and_GetArticleOutputPort_Success_and_return_nil",
			args: testData1.args,
			mocks: mocks{
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				authService: mocks_authService{
					currentUser_returns_error: derrors.Errors.NotAuthorized.Err,
				},
			},
			wants: usecase.Success(GetArticleResult{
				Article:   testData1.wants.Success().Article,
				Favorited: false,
				Favorites: testData1.wants.Success().Favorites,
			}),
		},
		{
			name: "when_given_not_exists_slug_should_call_ArticleRepositroy_Get_with_return_error_and_call_ErrorHandler_Handle_and_return_nil",
			args: args{
				slugStr: "notexists",
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 1,
					handle_returns_result:   failResult,
				},
				articleRepository: mocks_articleRepository{
					get_args_slugStr:    "notexists",
					get_returns_article: nil,
					get_returns_error:   dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call: true,
				favorite_listBySlug_should_skip: true,
				auth_currentUser_should_skip:    true,
			},
		},
		{
			name: "when_given_invalid_slug_should_call_ErrorHandler_and_return_nil",
			args: args{
				slugStr: strings.Repeat("a", 256),
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       derrors.Errors.Validation.Err,
					handle_args_opts_length: 1,
					handle_returns_result:   failResult,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call: true,
				article_get_should_skip:         true,
				favorite_listBySlug_should_skip: true,
				auth_currentUser_should_skip:    true,
			},
		},
		{
			name: "when_favorite_listBySlug_return_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   failResult,
				},
				articleRepository: testData1.mocks.articleRepository,
				favoriteRepository: mocks_favoriteRepository{
					listBySlug_args_slug:     testData1.mocks.favoriteRepository.listBySlug_args_slug,
					listBySlug_returns_error: dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call: true,
				auth_currentUser_should_skip:    true,
			},
		},
		{
			name: "when_authService_currentUser_return_unknown_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   failResult,
				},
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				authService: mocks_authService{
					currentUser_returns_error: dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorHandler := mock_usecase.NewMockErrorHandler[GetArticleResult](t)
			article := mock_repository.NewMockArticle(t)
			favorite := mock_repository.NewMockFavorite(t)
			authService := mock_auth_service.NewMockAuth(t)

			if tt.configs.errorHandler_handle_should_call {
				errorHandlerExpectations(t, errorHandler, errorHandlerExpectationsOption[GetArticleResult]{
					HandleArgsError:      tt.mocks.errorHandler.handle_args_error,
					HandleArgsOptsLength: tt.mocks.errorHandler.handle_args_opts_length,
					HandleReturnsResult:  tt.mocks.errorHandler.handle_returns_result,
				})
			}
			if !tt.configs.article_get_should_skip {
				article.EXPECT().
					Get(
						mock.Anything, tt.mocks.articleRepository.get_args_slugStr,
					).
					Return(
						tt.mocks.articleRepository.get_returns_article,
						tt.mocks.articleRepository.get_returns_error,
					)
			}
			if !tt.configs.favorite_listBySlug_should_skip {
				favorite.EXPECT().
					ListBySlug(
						mock.Anything,
						tt.mocks.favoriteRepository.listBySlug_args_slug,
					).
					Return(
						tt.mocks.favoriteRepository.listBySlug_returns_favorites,
						tt.mocks.favoriteRepository.listBySlug_returns_error,
					)
			}
			if !tt.configs.auth_currentUser_should_skip {
				authService.EXPECT().
					CurrentUser(mock.Anything).
					Return(
						tt.mocks.authService.currentUser_returns_user,
						tt.mocks.authService.currentUser_returns_error,
					)
			}

			a := NewGetArticle(errorHandler, article, favorite, authService)
			result := a.Get(context.Background(), tt.args.slugStr)
			assert.Equal(t, tt.wants, result)
		})
	}
}
