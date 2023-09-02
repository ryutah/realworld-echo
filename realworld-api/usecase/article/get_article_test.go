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
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetArticle_Get(t *testing.T) {
	type mocks_errorHandler struct {
		handle_args_error     error
		handle_args_opts      []usecase.ErrorHandlerOption
		handle_returns_result *usecase.Result[GetArticleResult]
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
	type mocks_followRepository struct {
		existsList_args_followedBy            authmodel.UserID
		existsList_args_following             []authmodel.UserID
		existsList_returns_followersExistsMap model.FollowersExistsMap
		existsList_returns_error              error
	}
	type mocks_authService struct {
		currentUser_returns_user  *authmodel.User
		currentUser_returns_error error
	}
	type mocks struct {
		errorHandler       mocks_errorHandler
		articleRepository  mocks_articleRepository
		favoriteRepository mocks_favoriteRepository
		followRepository   mocks_followRepository
		authService        mocks_authService
	}
	type configs struct {
		errorHandler_handle_should_call       bool
		article_get_should_be_skipped         bool
		favorite_listBySlug_should_be_skipped bool
		follow_existsList_should_be_skipped   bool
		auth_currentUser_should_be_skipped    bool
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
		author     authmodel.UserID = "Author"
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
						Author:    author,
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
				followRepository: mocks_followRepository{
					existsList_args_followedBy: user1,
					existsList_args_following:  []authmodel.UserID{author},
					existsList_returns_followersExistsMap: model.FollowersExistsMap{
						author: true,
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
					Author:    author,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Favorited: true,
				Favorites: model.FavoriteSlice{
					{ArticleSlug: slug1, UserID: user1},
					{ArticleSlug: slug1, UserID: user2},
				},
				FollowingAuthor: true,
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
			name:  "when_given_any_slug_should_return_success_result",
			args:  testData1.args,
			mocks: testData1.mocks,
			wants: testData1.wants,
		},
		{
			name: "when_given_any_slug_without_signedIn_user_should_return_nil",
			args: testData1.args,
			mocks: mocks{
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				authService: mocks_authService{
					currentUser_returns_error: derrors.Errors.NotAuthorized.Err,
				},
			},
			wants: usecase.Success(GetArticleResult{
				Article:         testData1.wants.Success().Article,
				Favorited:       false,
				Favorites:       testData1.wants.Success().Favorites,
				FollowingAuthor: false,
			}),
			configs: configs{
				follow_existsList_should_be_skipped: true,
			},
		},
		{
			name: "when_given_any_slug_with_ArticleRepositroy_Get_returns_error_should_call_ErrorHandler_and_return_fail_result",
			args: args{
				slugStr: "notexists",
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error: dummyError,
					handle_args_opts: []usecase.ErrorHandlerOption{
						usecase.WithNotFoundHandler(derrors.Errors.NotFound.Err),
					},
					handle_returns_result: failResult,
				},
				articleRepository: mocks_articleRepository{
					get_args_slugStr:  "notexists",
					get_returns_error: dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call:       true,
				favorite_listBySlug_should_be_skipped: true,
				follow_existsList_should_be_skipped:   true,
				auth_currentUser_should_be_skipped:    true,
			},
		},
		{
			name: "when_given_invalid_slug_should_call_ErrorHandler_and_return_fail_result",
			args: args{
				slugStr: strings.Repeat("a", 256),
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error: derrors.Errors.Validation.Err,
					handle_args_opts: []usecase.ErrorHandlerOption{
						usecase.WithBadRequestHandler(derrors.Errors.Validation.Err),
					},
					handle_returns_result: failResult,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call:       true,
				article_get_should_be_skipped:         true,
				favorite_listBySlug_should_be_skipped: true,
				follow_existsList_should_be_skipped:   true,
				auth_currentUser_should_be_skipped:    true,
			},
		},
		{
			name: "when_favorite_listBySlug_return_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: failResult,
				},
				articleRepository: testData1.mocks.articleRepository,
				favoriteRepository: mocks_favoriteRepository{
					listBySlug_args_slug:     testData1.mocks.favoriteRepository.listBySlug_args_slug,
					listBySlug_returns_error: dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call:     true,
				follow_existsList_should_be_skipped: true,
				auth_currentUser_should_be_skipped:  true,
			},
		},
		{
			name: "when_authService_currentUser_return_unknown_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: failResult,
				},
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				authService: mocks_authService{
					currentUser_returns_error: dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call:     true,
				follow_existsList_should_be_skipped: true,
			},
		},
		{
			name: "when_followRepo_existsList_return_unknown_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: failResult,
				},
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				followRepository: mocks_followRepository{
					existsList_args_followedBy: testData1.mocks.followRepository.existsList_args_followedBy,
					existsList_args_following:  testData1.mocks.followRepository.existsList_args_following,
					existsList_returns_error:   dummyError,
				},
				authService: testData1.mocks.authService,
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
			follow := mock_repository.NewMockFollow(t)
			authService := mock_auth_service.NewMockAuth(t)

			if tt.configs.errorHandler_handle_should_call {
				errorHandlerExpectations(t, errorHandler, errorHandlerExpectationsOption[GetArticleResult]{
					HandleArgsError:     tt.mocks.errorHandler.handle_args_error,
					HandleArgsOpts:      tt.mocks.errorHandler.handle_args_opts,
					HandleReturnsResult: tt.mocks.errorHandler.handle_returns_result,
				})
			}
			if !tt.configs.article_get_should_be_skipped {
				article.EXPECT().
					Get(
						mock.Anything, tt.mocks.articleRepository.get_args_slugStr,
					).
					Return(
						tt.mocks.articleRepository.get_returns_article,
						tt.mocks.articleRepository.get_returns_error,
					)
			}
			if !tt.configs.favorite_listBySlug_should_be_skipped {
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
			if !tt.configs.auth_currentUser_should_be_skipped {
				authService.EXPECT().
					CurrentUser(mock.Anything).
					Return(
						tt.mocks.authService.currentUser_returns_user,
						tt.mocks.authService.currentUser_returns_error,
					)
			}
			if !tt.configs.follow_existsList_should_be_skipped {
				follow.EXPECT().
					ExistsList(
						mock.Anything,
						tt.mocks.followRepository.existsList_args_followedBy,
						lo.ToAnySlice(tt.mocks.followRepository.existsList_args_following)...,
					).
					Return(
						tt.mocks.followRepository.existsList_returns_followersExistsMap,
						tt.mocks.followRepository.existsList_returns_error,
					)
			}

			a := NewGetArticle(errorHandler, article, favorite, follow, authService)
			result := a.Get(context.Background(), tt.args.slugStr)
			assert.Equal(t, tt.wants, result)
		})
	}
}
