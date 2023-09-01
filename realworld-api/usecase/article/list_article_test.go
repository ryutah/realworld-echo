package article_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	mock_repository "github.com/ryutah/realworld-echo/realworld-api/internal/mock/article/repository"
	mock_auth_service "github.com/ryutah/realworld-echo/realworld-api/internal/mock/auth/service"
	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/pointer"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ListArticle_List(t *testing.T) {
	type args struct {
		param ListArticleParam
	}
	type mocks_errorHandler struct {
		handle_args_error       error
		handle_args_opts_length int
		handle_returns_result   *usecase.Result[ListArticleResult]
	}
	type mocks_articleRepository struct {
		search_args_articleSearchParam repository.ArticleSearchParam
		search_retunrs_articles        []model.Article
		search_results_error           error
	}
	type mocks_favoriteRepository struct {
		listBySlugs_args_slugs        []model.Slug
		listBySlugs_returns_favorites model.FavoriteSliceMap
		listBySlugs_returns_error     error
	}
	type mocks_authService struct {
		currentUser_returns_user *authmodel.User
		currentUser_returns_err  error
	}

	type mocks struct {
		errorHandler       mocks_errorHandler
		articleRepository  mocks_articleRepository
		favoriteRepository mocks_favoriteRepository
		authService        mocks_authService
	}
	type wants struct {
		result *usecase.Result[ListArticleResult]
	}
	type configs struct {
		errorHandler_handle_should_call  bool
		article_search_should_skip       bool
		favorite_listBySlugs_should_skip bool
		auth_currentUser_should_skip     bool
	}

	var (
		tag, tErr                                           = model.NewTagName("tag")
		dummyError                                          = errors.New("dummy")
		badrequestResult *usecase.Result[ListArticleResult] = usecase.Fail[ListArticleResult](
			usecase.NewFailResult(usecase.FailTypeBadRequest, "fail"),
		)
		internalErrorResult *usecase.Result[ListArticleResult] = usecase.Fail[ListArticleResult](
			usecase.NewFailResult(usecase.FailTypeInternalError, "fail"),
		)
		slug1, _    = model.NewSlug(uuid.New().String())
		slug2, _    = model.NewSlug(uuid.New().String())
		user1       = authmodel.UserID("user1")
		user2       = authmodel.UserID("user2")
		favoritedBy = authmodel.UserID("user2")
		testData1   = struct {
			args  args
			mocks mocks
			wants wants
		}{
			args: args{
				param: ListArticleParam{
					Tag:         "tag",
					Author:      user1.String(),
					FavoritedBy: favoritedBy.String(),
					Offset:      10,
					Limit:       20,
				},
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Tag:         &tag,
						Author:      &user1,
						FavoritedBy: &favoritedBy,
						Offset:      10,
						Limit:       20,
					},
					search_retunrs_articles: []model.Article{
						{Slug: slug1, Author: user1},
						{Slug: slug2, Author: user2},
					},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs: []model.Slug{
						slug1, slug2,
					},
					listBySlugs_returns_favorites: model.FavoriteSliceMap{
						slug1: model.FavoriteSlice{
							{ArticleSlug: slug1, UserID: user1},
							{ArticleSlug: slug2, UserID: user2},
						},
						slug2: model.FavoriteSlice{
							{ArticleSlug: slug2, UserID: user2},
						},
					},
				},
				authService: mocks_authService{
					currentUser_returns_user: &authmodel.User{ID: user1},
				},
			},
			wants: wants{
				result: usecase.Success[ListArticleResult](ListArticleResult{
					Articles: []ListArticleResultArtile{
						{
							Aritcle: model.Article{Slug: slug1, Author: user1},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: slug1, UserID: user1},
								{ArticleSlug: slug2, UserID: user2},
							},
							Favorited: true,
						},
						{
							Aritcle: model.Article{Slug: slug2, Author: user2},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: slug2, UserID: user2},
							},
							Favorited: false,
						},
					},
				}),
			},
		}
	)
	if tErr != nil {
		t.Fatal(tErr)
	}

	tests := []struct {
		name    string
		args    args
		mocks   mocks
		wants   wants
		configs configs
	}{
		{
			name:    "valid_params_should_returns_expected_result",
			args:    testData1.args,
			mocks:   testData1.mocks,
			wants:   testData1.wants,
			configs: configs{},
		},
		{
			name: "not_authorized_user_valid_params_should_returns_expected_result",
			args: testData1.args,
			mocks: mocks{
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				authService: mocks_authService{
					currentUser_returns_err: derrors.Errors.NotAuthorized.Err,
				},
			},
			wants: wants{
				result: usecase.Success[ListArticleResult](ListArticleResult{
					Articles: lo.Map(
						testData1.wants.result.Success().Articles,
						func(a ListArticleResultArtile, _ int) ListArticleResultArtile {
							return ListArticleResultArtile{
								Aritcle:   a.Aritcle,
								Favorites: a.Favorites,
								Favorited: false,
							}
						},
					),
				}),
			},
			configs: configs{},
		},
		{
			name: "valid_params_with_zero_result_should_returns_expected_result",
			args: testData1.args,
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: testData1.mocks.articleRepository.search_args_articleSearchParam,
					search_retunrs_articles:        []model.Article{},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs:        []model.Slug{},
					listBySlugs_returns_favorites: model.FavoriteSliceMap{},
				},
				authService: testData1.mocks.authService,
			},
			wants: wants{
				result: usecase.Success[ListArticleResult](ListArticleResult{
					Articles: []ListArticleResultArtile{},
				}),
			},
			configs: configs{},
		},
		{
			name: "invalid_params_should_returns_validation_error",
			args: args{
				param: ListArticleParam{
					Tag: strings.Repeat("a", 10000),
				},
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       derrors.Errors.Validation.Err,
					handle_args_opts_length: 1,
					handle_returns_result:   badrequestResult,
				},
			},
			wants: wants{
				result: badrequestResult,
			},
			configs: configs{
				errorHandler_handle_should_call:  true,
				article_search_should_skip:       true,
				favorite_listBySlugs_should_skip: true,
				auth_currentUser_should_skip:     true,
			},
		},
		{
			name: "article_repository_search_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorResult,
				},
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: testData1.mocks.articleRepository.search_args_articleSearchParam,
					search_results_error:           dummyError,
				},
			},
			wants: wants{
				result: internalErrorResult,
			},
			configs: configs{
				errorHandler_handle_should_call:  true,
				favorite_listBySlugs_should_skip: true,
				auth_currentUser_should_skip:     true,
			},
		},
		{
			name: "favorite_repository_listBySlugs_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorResult,
				},
				articleRepository: testData1.mocks.articleRepository,
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs:    testData1.mocks.favoriteRepository.listBySlugs_args_slugs,
					listBySlugs_returns_error: dummyError,
				},
			},
			wants: wants{
				result: internalErrorResult,
			},
			configs: configs{
				errorHandler_handle_should_call: true,
				auth_currentUser_should_skip:    true,
			},
		},
		{
			name: "authService_currentUser_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorResult,
				},
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				authService: mocks_authService{
					currentUser_returns_err: dummyError,
				},
			},
			wants: wants{
				result: internalErrorResult,
			},
			configs: configs{
				errorHandler_handle_should_call: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorHandler := mock_usecase.NewMockErrorHandler[ListArticleResult](t)
			articleRepository := mock_repository.NewMockArticle(t)
			favoriteRepository := mock_repository.NewMockFavorite(t)
			authService := mock_auth_service.NewMockAuth(t)

			if tt.configs.errorHandler_handle_should_call {
				errorHandlerExpectations(t, errorHandler, errorHandlerExpectationsOption[ListArticleResult]{
					HandleArgsError:      tt.mocks.errorHandler.handle_args_error,
					HandleArgsOptsLength: tt.mocks.errorHandler.handle_args_opts_length,
					HandleReturnsResult:  tt.mocks.errorHandler.handle_returns_result,
				})
			}
			if !tt.configs.article_search_should_skip {
				articleRepository.EXPECT().
					Search(
						mock.Anything, tt.mocks.articleRepository.search_args_articleSearchParam,
					).
					Return(
						tt.mocks.articleRepository.search_retunrs_articles,
						tt.mocks.articleRepository.search_results_error,
					)
			}
			if !tt.configs.favorite_listBySlugs_should_skip {
				favoriteRepository.EXPECT().
					ListBySlugs(
						mock.Anything,
						lo.ToAnySlice(tt.mocks.favoriteRepository.listBySlugs_args_slugs)...,
					).
					Return(
						tt.mocks.favoriteRepository.listBySlugs_returns_favorites,
						tt.mocks.favoriteRepository.listBySlugs_returns_error,
					)
			}
			if !tt.configs.auth_currentUser_should_skip {
				authService.EXPECT().
					CurrentUser(mock.Anything).
					Return(
						tt.mocks.authService.currentUser_returns_user,
						tt.mocks.authService.currentUser_returns_err,
					)
			}

			a := NewListArticle[any](errorHandler, articleRepository, favoriteRepository, authService)
			got := a.List(context.Background(), tt.args.param)
			assert.Equal(t, tt.wants.result, got)
		})
	}
}

func TestListArticleParam_ToSearchParam(t *testing.T) {
	type fields struct {
		Tag         string
		Author      string
		FavoritedBy string
	}
	type wants struct {
		param *repository.ArticleSearchParam
		err   error
	}

	tests := []struct {
		name   string
		fields fields
		target ListArticleParam
		want   wants
	}{
		{
			name: "valid_params_should_returns_expected_result",
			target: ListArticleParam{
				Tag:         "tag",
				Author:      "author",
				FavoritedBy: "favorited_by",
				Offset:      10,
				Limit:       20,
			},
			want: wants{
				param: &repository.ArticleSearchParam{
					Tag:         mustNewTagName("tag"),
					Author:      pointer.Pointer[authmodel.UserID]("author"),
					FavoritedBy: pointer.Pointer[authmodel.UserID]("favorited_by"),
					Offset:      10,
					Limit:       20,
				},
				err: nil,
			},
		},
		{
			name:   "blank_params_should_returns_expected_result",
			target: ListArticleParam{},
			want: wants{
				param: &repository.ArticleSearchParam{
					Limit: repository.DefaultLimit,
				},
				err: nil,
			},
		},
		{
			name: "invalid_tag_should_returns_validation_error",
			target: ListArticleParam{
				Tag: strings.Repeat("a", 10000),
			},
			want: wants{
				err: derrors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_author_should_returns_validation_error",
			target: ListArticleParam{
				Author: strings.Repeat("a", 10000),
			},
			want: wants{
				err: derrors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_favarited_by_should_returns_validation_error",
			target: ListArticleParam{
				FavoritedBy: strings.Repeat("a", 10000),
			},
			want: wants{
				err: derrors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_limit_should_returns_validation_error",
			target: ListArticleParam{
				Limit: 1000000,
			},
			want: wants{
				err: derrors.Errors.Validation.Err,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.target.ToSearchParam()

			assert.Equal(t, tt.want.param, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
