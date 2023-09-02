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
		handle_args_error     error
		handle_args_opts      []usecase.ErrorHandlerOption
		handle_returns_result *usecase.Result[ListArticleResult]
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
	type mocks_followRepository struct {
		existsList_args_followedBy            authmodel.UserID
		existsList_args_followings            []authmodel.UserID
		existsList_returns_followersExistsMap model.FollowersExistsMap
		existsList_returns_error              error
	}
	type mocks_authService struct {
		currentUser_returns_user *authmodel.User
		currentUser_returns_err  error
	}
	type mocks struct {
		errorHandler       mocks_errorHandler
		articleRepository  mocks_articleRepository
		favoriteRepository mocks_favoriteRepository
		followRepository   mocks_followRepository
		authService        mocks_authService
	}
	type wants struct {
		result *usecase.Result[ListArticleResult]
	}
	type configs struct {
		errorHandler_handle_should_be_called   bool
		article_search_should_be_skipped       bool
		favorite_listBySlugs_should_be_skipped bool
		follow_listBySlugs_should_be_skipped   bool
		auth_currentUser_should_be_skipped     bool
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
		slug1, _  = model.NewSlug(uuid.New().String())
		slug2, _  = model.NewSlug(uuid.New().String())
		user1     = authmodel.UserID("user1")
		user2     = authmodel.UserID("user2")
		author1   = authmodel.UserID("author1")
		author2   = authmodel.UserID("author2")
		testData1 = struct {
			args  args
			mocks mocks
			wants wants
		}{
			args: args{
				param: ListArticleParam{
					Tag:         "tag",
					Author:      author1.String(),
					FavoritedBy: user1.String(),
					Offset:      10,
					Limit:       20,
				},
			},
			mocks: mocks{
				articleRepository: mocks_articleRepository{
					search_args_articleSearchParam: repository.ArticleSearchParam{
						Tag:         &tag,
						Author:      &author1,
						FavoritedBy: &user1,
						Offset:      10,
						Limit:       20,
					},
					search_retunrs_articles: []model.Article{
						{Slug: slug1, Author: author1},
						{Slug: slug2, Author: author2},
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
							{ArticleSlug: slug2, UserID: user1},
						},
					},
				},
				followRepository: mocks_followRepository{
					existsList_args_followedBy: user1,
					existsList_args_followings: []authmodel.UserID{
						author1, author2,
					},
					existsList_returns_followersExistsMap: model.FollowersExistsMap{
						author1: true,
						author2: false,
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
							Aritcle: model.Article{Slug: slug1, Author: author1},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: slug1, UserID: user1},
								{ArticleSlug: slug2, UserID: user2},
							},
							Favorited:       true,
							AuthorFollowing: true,
						},
						{
							Aritcle: model.Article{Slug: slug2, Author: author2},
							Favorites: model.FavoriteSlice{
								{ArticleSlug: slug2, UserID: user1},
							},
							Favorited:       true,
							AuthorFollowing: false,
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
			configs: configs{
				follow_listBySlugs_should_be_skipped: true,
			},
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
				followRepository: mocks_followRepository{
					existsList_args_followedBy:            testData1.mocks.followRepository.existsList_args_followedBy,
					existsList_args_followings:            []authmodel.UserID{},
					existsList_returns_followersExistsMap: model.FollowersExistsMap{},
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
					handle_args_error: derrors.Errors.Validation.Err,
					handle_args_opts: []usecase.ErrorHandlerOption{
						usecase.WithBadRequestHandler(derrors.Errors.Validation.Err),
					},
					handle_returns_result: badrequestResult,
				},
			},
			wants: wants{
				result: badrequestResult,
			},
			configs: configs{
				errorHandler_handle_should_be_called:   true,
				article_search_should_be_skipped:       true,
				favorite_listBySlugs_should_be_skipped: true,
				follow_listBySlugs_should_be_skipped:   true,
				auth_currentUser_should_be_skipped:     true,
			},
		},
		{
			name: "article_repository_search_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: internalErrorResult,
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
				errorHandler_handle_should_be_called:   true,
				favorite_listBySlugs_should_be_skipped: true,
				follow_listBySlugs_should_be_skipped:   true,
				auth_currentUser_should_be_skipped:     true,
			},
		},
		{
			name: "favorite_repository_listBySlugs_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: internalErrorResult,
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
				errorHandler_handle_should_be_called: true,
				follow_listBySlugs_should_be_skipped: true,
				auth_currentUser_should_be_skipped:   true,
			},
		},
		{
			name: "authService_currentUser_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: internalErrorResult,
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
				errorHandler_handle_should_be_called: true,
				follow_listBySlugs_should_be_skipped: true,
			},
		},
		{
			name: "followRepo_existsList_failed_should_returns_failed_result",
			args: testData1.args,
			mocks: mocks{
				errorHandler: mocks_errorHandler{
					handle_args_error:     dummyError,
					handle_returns_result: internalErrorResult,
				},
				articleRepository:  testData1.mocks.articleRepository,
				favoriteRepository: testData1.mocks.favoriteRepository,
				followRepository: mocks_followRepository{
					existsList_args_followedBy: testData1.mocks.followRepository.existsList_args_followedBy,
					existsList_args_followings: testData1.mocks.followRepository.existsList_args_followings,
					existsList_returns_error:   dummyError,
				},
				authService: testData1.mocks.authService,
			},
			wants: wants{
				result: internalErrorResult,
			},
			configs: configs{
				errorHandler_handle_should_be_called: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorHandler := mock_usecase.NewMockErrorHandler[ListArticleResult](t)
			articleRepository := mock_repository.NewMockArticle(t)
			favoriteRepository := mock_repository.NewMockFavorite(t)
			followRepository := mock_repository.NewMockFollow(t)
			authService := mock_auth_service.NewMockAuth(t)

			if tt.configs.errorHandler_handle_should_be_called {
				errorHandlerExpectations(t, errorHandler, errorHandlerExpectationsOption[ListArticleResult]{
					HandleArgsError:     tt.mocks.errorHandler.handle_args_error,
					HandleArgsOpts:      tt.mocks.errorHandler.handle_args_opts,
					HandleReturnsResult: tt.mocks.errorHandler.handle_returns_result,
				})
			}
			if !tt.configs.article_search_should_be_skipped {
				articleRepository.EXPECT().
					Search(
						mock.Anything, tt.mocks.articleRepository.search_args_articleSearchParam,
					).
					Return(
						tt.mocks.articleRepository.search_retunrs_articles,
						tt.mocks.articleRepository.search_results_error,
					)
			}
			if !tt.configs.favorite_listBySlugs_should_be_skipped {
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
			if !tt.configs.auth_currentUser_should_be_skipped {
				authService.EXPECT().
					CurrentUser(mock.Anything).
					Return(
						tt.mocks.authService.currentUser_returns_user,
						tt.mocks.authService.currentUser_returns_err,
					)
			}
			if !tt.configs.follow_listBySlugs_should_be_skipped {
				followRepository.EXPECT().
					ExistsList(
						mock.Anything,
						tt.mocks.followRepository.existsList_args_followedBy,
						lo.ToAnySlice(tt.mocks.followRepository.existsList_args_followings)...,
					).
					Return(
						tt.mocks.followRepository.existsList_returns_followersExistsMap,
						tt.mocks.followRepository.existsList_returns_error,
					)
			}

			a := NewListArticle(errorHandler, articleRepository, favoriteRepository, followRepository, authService)
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
