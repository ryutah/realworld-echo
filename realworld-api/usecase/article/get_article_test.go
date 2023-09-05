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
	mock_article_repo "github.com/ryutah/realworld-echo/realworld-api/internal/mock/gen/domain/article/repository"
	mock_auth_service "github.com/ryutah/realworld-echo/realworld-api/internal/mock/gen/domain/auth/service"
	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/gen/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase/article"
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
		count_args_slug       model.Slug
		count_returns_count   int
		count_returns_error   error
		exists_args_user      authmodel.UserID
		exists_args_slug      model.Slug
		exists_returns_exists bool
		exists_returns_error  error
	}
	type mocks_followRepository struct {
		exists_args_followedBy authmodel.UserID
		exists_args_following  authmodel.UserID
		exists_returns_exists  bool
		exists_returns_error   error
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
		errorHandler_handle_should_call    bool
		article_get_should_be_skipped      bool
		favorite_count_should_be_skipped   bool
		favorite_exists_should_be_skipped  bool
		follow_exists_should_be_skipped    bool
		auth_currentUser_should_be_skipped bool
	}
	type args struct {
		slugStr string
	}

	var (
		now                         = premitive.NewJSTTime(time.Now())
		dummyError                  = errors.New("dummy")
		slug1      model.Slug       = "slug"
		user1      authmodel.UserID = "user1"
		author                      = model.UserProfile{ID: "Author"}
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
					count_args_slug:       slug1,
					count_returns_count:   2,
					exists_args_user:      user1,
					exists_args_slug:      slug1,
					exists_returns_exists: true,
				},
				followRepository: mocks_followRepository{
					exists_args_followedBy: user1,
					exists_args_following:  author.ID,
					exists_returns_exists:  true,
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
				Favorited:       true,
				FavoriteCount:   2,
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
				FavoriteCount:   testData1.wants.Success().FavoriteCount,
				FollowingAuthor: false,
			}),
			configs: configs{
				favorite_exists_should_be_skipped: true,
				follow_exists_should_be_skipped:   true,
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
				errorHandler_handle_should_call:    true,
				favorite_count_should_be_skipped:   true,
				favorite_exists_should_be_skipped:  true,
				follow_exists_should_be_skipped:    true,
				auth_currentUser_should_be_skipped: true,
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
				errorHandler_handle_should_call:    true,
				article_get_should_be_skipped:      true,
				favorite_count_should_be_skipped:   true,
				favorite_exists_should_be_skipped:  true,
				follow_exists_should_be_skipped:    true,
				auth_currentUser_should_be_skipped: true,
			},
		},
		{
			name: "when_favorite_count_return_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: failResult,
				},
				articleRepository: testData1.mocks.articleRepository,
				favoriteRepository: mocks_favoriteRepository{
					count_args_slug:     testData1.mocks.favoriteRepository.count_args_slug,
					count_returns_error: dummyError,
				},
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call:    true,
				favorite_exists_should_be_skipped:  true,
				follow_exists_should_be_skipped:    true,
				auth_currentUser_should_be_skipped: true,
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
				errorHandler_handle_should_call:   true,
				favorite_exists_should_be_skipped: true,
				follow_exists_should_be_skipped:   true,
			},
		},
		{
			name: "when_favoriteRepo_exists_return_unknown_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: failResult,
				},
				articleRepository: testData1.mocks.articleRepository,
				favoriteRepository: mocks_favoriteRepository{
					count_args_slug:      testData1.mocks.favoriteRepository.count_args_slug,
					count_returns_count:  testData1.mocks.favoriteRepository.count_returns_count,
					exists_args_user:     testData1.mocks.favoriteRepository.exists_args_user,
					exists_args_slug:     testData1.mocks.favoriteRepository.exists_args_slug,
					exists_returns_error: dummyError,
				},
				authService: testData1.mocks.authService,
			},
			wants: failResult,
			configs: configs{
				errorHandler_handle_should_call: true,
				follow_exists_should_be_skipped: true,
			},
		},
		{
			name: "when_followRepo_exists_return_unknown_error_should_call_ErrorHandler_and_return_fail_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: failResult,
				},
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				followRepository: mocks_followRepository{
					exists_args_followedBy: testData1.mocks.followRepository.exists_args_followedBy,
					exists_args_following:  testData1.mocks.followRepository.exists_args_following,
					exists_returns_error:   dummyError,
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
			article := mock_article_repo.NewMockArticle(t)
			favorite := mock_article_repo.NewMockFavorite(t)
			follow := mock_article_repo.NewMockFollow(t)
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
			if !tt.configs.favorite_count_should_be_skipped {
				favorite.EXPECT().
					Count(
						mock.Anything,
						tt.mocks.favoriteRepository.count_args_slug,
					).
					Return(
						tt.mocks.favoriteRepository.count_returns_count,
						tt.mocks.favoriteRepository.count_returns_error,
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
			if !tt.configs.favorite_exists_should_be_skipped {
				favorite.EXPECT().
					Exists(
						mock.Anything,
						tt.mocks.favoriteRepository.exists_args_user,
						tt.mocks.favoriteRepository.exists_args_slug,
					).
					Return(
						tt.mocks.favoriteRepository.exists_returns_exists,
						tt.mocks.favoriteRepository.exists_returns_error,
					)
			}
			if !tt.configs.follow_exists_should_be_skipped {
				follow.EXPECT().
					Exists(
						mock.Anything,
						tt.mocks.followRepository.exists_args_followedBy,
						tt.mocks.followRepository.exists_args_following,
					).
					Return(
						tt.mocks.followRepository.exists_returns_exists,
						tt.mocks.followRepository.exists_returns_error,
					)
			}

			a := NewGetArticle(errorHandler, article, favorite, follow, authService)
			result := a.Get(context.Background(), tt.args.slugStr)
			assert.Equal(t, tt.wants, result)
		})
	}
}
