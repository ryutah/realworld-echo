package article_test

import (
	"context"
	"errors"
	"strings"
	"testing"

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
		ctx   context.Context
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
		tag, tErr                                           = model.NewArticleTag("tag")
		authID           authmodel.UserID                   = "author"
		favoritedBy      authmodel.UserID                   = "favorited_by"
		dummyError                                          = errors.New("dummy")
		badrequestResult *usecase.Result[ListArticleResult] = usecase.Fail[ListArticleResult](
			usecase.NewFailResult(usecase.FailTypeBadRequest, "fail"),
		)
		internalErrorResult *usecase.Result[ListArticleResult] = usecase.Fail[ListArticleResult](
			usecase.NewFailResult(usecase.FailTypeInternalError, "fail"),
		)
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
			name: "valid_params_should_returns_expected_result",
			args: args{
				ctx: context.TODO(),
				param: ListArticleParam{
					Tag:         "tag",
					Author:      "author",
					FavoritedBy: "favorited_by",
					Offset:      10,
					Limit:       20,
				},
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Tag:         tag,
						Author:      &authID,
						FavoritedBy: &favoritedBy,
						Offset:      10,
						Limit:       20,
					},
					search_retunrs_articles: []model.Article{
						{Slug: "dummy"},
						{Slug: "dummy2"},
					},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs: []model.Slug{
						"dummy", "dummy2",
					},
					listBySlugs_returns_favorites: model.FavoriteSliceMap{
						"dummy": model.FavoriteSlice{
							{ArticleSlug: "dummy", UserID: "user1"},
							{ArticleSlug: "dummy", UserID: "user2"},
						},
						"dummy2": model.FavoriteSlice{
							{ArticleSlug: "dummy2", UserID: "user2"},
						},
					},
				},
				authService: mocks_authService{
					currentUser_returns_user: &authmodel.User{ID: "user1"},
				},
			},
			wants: wants{
				result: usecase.Success[ListArticleResult](ListArticleResult{
					Articles: []ListArticleResultArtile{
						{
							Aritcle: model.Article{Slug: "dummy"},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: "dummy", UserID: "user1"},
								{ArticleSlug: "dummy", UserID: "user2"},
							},
							Favorited: true,
						},
						{
							Aritcle: model.Article{Slug: "dummy2"},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: "dummy2", UserID: "user2"},
							},
							Favorited: false,
						},
					},
				}),
			},
			configs: configs{},
		},
		{
			name: "not_authorized_user_valid_params_should_returns_expected_result",
			args: args{
				ctx: context.TODO(),
				param: ListArticleParam{
					Tag:         "tag",
					Author:      "author",
					FavoritedBy: "favorited_by",
					Offset:      10,
					Limit:       20,
				},
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Tag:         tag,
						Author:      &authID,
						FavoritedBy: &favoritedBy,
						Offset:      10,
						Limit:       20,
					},
					search_retunrs_articles: []model.Article{
						{Slug: "dummy"},
						{Slug: "dummy2"},
					},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs: []model.Slug{
						"dummy", "dummy2",
					},
					listBySlugs_returns_favorites: model.FavoriteSliceMap{
						"dummy": model.FavoriteSlice{
							{ArticleSlug: "dummy", UserID: "user1"},
							{ArticleSlug: "dummy", UserID: "user2"},
						},
						"dummy2": model.FavoriteSlice{
							{ArticleSlug: "dummy2", UserID: "user2"},
						},
					},
				},
				authService: mocks_authService{
					currentUser_returns_err: derrors.Errors.NotAuthorized.Err,
				},
			},
			wants: wants{
				result: usecase.Success[ListArticleResult](ListArticleResult{
					Articles: []ListArticleResultArtile{
						{
							Aritcle: model.Article{Slug: "dummy"},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: "dummy", UserID: "user1"},
								{ArticleSlug: "dummy", UserID: "user2"},
							},
							Favorited: false,
						},
						{
							Aritcle: model.Article{Slug: "dummy2"},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: "dummy2", UserID: "user2"},
							},
							Favorited: false,
						},
					},
				}),
			},
			configs: configs{},
		},
		{
			name: "valid_params_with_zero_result_should_returns_expected_result",
			args: args{
				ctx: context.TODO(),
				param: ListArticleParam{
					Tag:         "tag",
					Author:      "author",
					FavoritedBy: "favorited_by",
					Offset:      10,
					Limit:       20,
				},
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Tag:         tag,
						Author:      &authID,
						FavoritedBy: &favoritedBy,
						Offset:      10,
						Limit:       20,
					},
					search_retunrs_articles: []model.Article{},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs:        []model.Slug{},
					listBySlugs_returns_favorites: model.FavoriteSliceMap{},
				},
				authService: mocks_authService{
					currentUser_returns_user: &authmodel.User{ID: "user1"},
				},
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
				ctx: context.TODO(),
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
			args: args{
				ctx:   context.TODO(),
				param: ListArticleParam{},
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorResult,
				},
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Limit: repository.DefaultLimit,
					},
					search_results_error: dummyError,
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
			args: args{
				ctx:   context.TODO(),
				param: ListArticleParam{},
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorResult,
				},
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Limit: repository.DefaultLimit,
					},
					search_retunrs_articles: []model.Article{
						{Slug: "dummy"},
					},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs: []model.Slug{
						"dummy",
					},
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
			args: args{
				ctx:   context.TODO(),
				param: ListArticleParam{},
			},
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorResult,
				},
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Limit: repository.DefaultLimit,
					},
					search_retunrs_articles: []model.Article{
						{Slug: "dummy"},
					},
				},
				favoriteRepository: mocks_favoriteRepository{
					listBySlugs_args_slugs: []model.Slug{
						"dummy",
					},
					listBySlugs_returns_favorites: model.FavoriteSliceMap{
						"dummy": model.FavoriteSlice{
							{ArticleSlug: "dummy", UserID: "user1"},
						},
					},
				},
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
				errorHandler.EXPECT().
					Handle(
						mock.Anything, mock.Anything, mock.Anything,
					).
					Run(func(ctx context.Context, err error, opts ...usecase.ErrorHandlerOption) {
						assert.ErrorIs(t, err, tt.mocks.errorHandler.handle_args_error, "error of ErrorHandler#Handle args")
						assert.Len(t, opts, tt.mocks.errorHandler.handle_args_opts_length, "length of ErrorHandler#Hanel option args")
					}).
					Return(
						tt.mocks.errorHandler.handle_returns_result,
					)
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
			got := a.List(tt.args.ctx, tt.args.param)
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
					Tag:         mustNewTag("tag"),
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
