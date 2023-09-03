package sqlc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	. "github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc/gen"
	mock_auth_repository "github.com/ryutah/realworld-echo/realworld-api/internal/mock/auth/repository"
	mock_psql_sqlc "github.com/ryutah/realworld-echo/realworld-api/internal/mock/psql/sqlc"
	mock_psql_sqlc_gen "github.com/ryutah/realworld-echo/realworld-api/internal/mock/psql/sqlc/gen"
	mock_transaction "github.com/ryutah/realworld-echo/realworld-api/internal/mock/transaction"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestArticle_GenerateID(t *testing.T) {
	type wants struct {
		emptySlug bool
		err       error
	}

	tests := []struct {
		name string
		want wants
	}{
		{
			name: "should_return_not_blank_slug",
			want: wants{
				emptySlug: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbManager := mock_psql_sqlc.NewMockDBManager(t)
			transaction := mock_transaction.NewMockTransaction(t)
			userRepo := mock_auth_repository.NewMockUser(t)
			articleSelector := mock_psql_sqlc.NewMockRawSelector[gen.Article](t)

			got, err := NewArtile(dbManager, transaction, userRepo, articleSelector).GenerateID(context.Background())
			if !tt.want.emptySlug {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
			if !assert.ErrorIs(t, err, tt.want.err) {
				t.Logf("%+v", err)
			}
		})
	}
}

func TestArticle_Get(t *testing.T) {
	type args struct {
		slug model.Slug
	}
	type mock_querier struct {
		getArticle_args_slug                uuid.UUID
		getArticle_returns_article          gen.Article
		getArticle_returns_error            error
		listArticleTags_args_slugs          []uuid.UUID
		listArticleTags_returns_articleTags []gen.ListArticleTagsRow
		listArticleTags_returns_error       error
	}
	type mock_userRepository struct {
		get_args_id       authmodel.UserID
		get_returns_user  *authmodel.User
		get_returns_error error
	}
	type mocks struct {
		querier        mock_querier
		userRepository mock_userRepository
	}
	type wants struct {
		article *model.Article
		err     error
	}
	type configs struct {
		dbManager_querier_should_be_skipped       bool
		querier_getArticle_should_be_skipped      bool
		querier_listArticleTags_should_be_skipped bool
		userRepo_get_should_be_skipped            bool
	}

	var (
		slug      = uuid.New()
		now       = time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC)
		nowTz     = ToTimestamptz(now)
		dummyErr  = errors.New("dummy")
		testData1 = struct {
			args  args
			mocks mocks
			wants wants
		}{
			args: args{
				slug: model.Slug(slug.String()),
			},
			mocks: mocks{
				querier: mock_querier{
					getArticle_args_slug: slug,
					getArticle_returns_article: gen.Article{
						Slug:        slug,
						Author:      "author",
						Body:        "body",
						Title:       "title",
						Description: "description",
						CreatedAt:   nowTz,
						UpdatedAt:   nowTz,
					},
					listArticleTags_args_slugs: []uuid.UUID{
						slug,
					},
					listArticleTags_returns_articleTags: []gen.ListArticleTagsRow{
						{
							ArticleSlug: slug,
							Name:        "tag1",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
						{
							ArticleSlug: slug,
							Name:        "tag2",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
					},
				},
				userRepository: mock_userRepository{
					get_args_id: authmodel.UserID("author"),
					get_returns_user: &authmodel.User{
						ID: "author",
						Account: authmodel.Account{
							Email: "sample@gmail.com",
						},
						Profile: authmodel.Profile{
							Username: "name",
							Bio:      "bio",
							Image:    "http://xxxxxxxx.png",
						},
					},
				},
			},
			wants: wants{
				article: &model.Article{
					Slug: model.Slug(slug.String()),
					Tags: []model.TagName{
						"tag1", "tag2",
					},
					Author: model.UserProfile{
						ID:    "author",
						Name:  "name",
						Bio:   "bio",
						Image: "http://xxxxxxxx.png",
					},
					Contents: model.ArticleContents{
						Title:       "title",
						Description: "description",
						Body:        "body",
					},
					CreatedAt: premitive.NewJSTTime(now),
					UpdatedAt: premitive.NewJSTTime(now),
				},
			},
		}
	)

	tests := []struct {
		name    string
		args    args
		mocks   mocks
		wants   wants
		configs configs
	}{
		{
			name:  "given_valid_slug_should_return_expected_article",
			args:  testData1.args,
			mocks: testData1.mocks,
			wants: testData1.wants,
		},
		{
			name: "given_valid_slug_not_exists_tags_should_return_expected_article",
			args: testData1.args,
			mocks: mocks{
				querier: mock_querier{
					getArticle_args_slug:                testData1.mocks.querier.getArticle_args_slug,
					getArticle_returns_article:          testData1.mocks.querier.getArticle_returns_article,
					listArticleTags_args_slugs:          testData1.mocks.querier.listArticleTags_args_slugs,
					listArticleTags_returns_articleTags: nil,
				},
				userRepository: testData1.mocks.userRepository,
			},
			wants: wants{
				article: &model.Article{
					Slug:      testData1.wants.article.Slug,
					Tags:      nil,
					Author:    testData1.wants.article.Author,
					Contents:  testData1.wants.article.Contents,
					CreatedAt: testData1.wants.article.CreatedAt,
					UpdatedAt: testData1.wants.article.UpdatedAt,
				},
			},
		},
		{
			name: "given_invalid_slug_should_return_validation_error",
			args: args{
				slug: model.Slug("invalid_uuid"),
			},
			mocks: mocks{},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
			configs: configs{
				dbManager_querier_should_be_skipped:       true,
				querier_getArticle_should_be_skipped:      true,
				querier_listArticleTags_should_be_skipped: true,
				userRepo_get_should_be_skipped:            true,
			},
		},
		{
			name: "given_valid_slug_with_querier_getArticle_returns_no_rows_error_should_return_not_found_error",
			args: testData1.args,
			mocks: mocks{
				querier: mock_querier{
					getArticle_args_slug:     testData1.mocks.querier.getArticle_args_slug,
					getArticle_returns_error: pgx.ErrNoRows,
				},
			},
			wants: wants{
				err: derrors.Errors.NotFound.Err,
			},
			configs: configs{
				querier_listArticleTags_should_be_skipped: true,
				userRepo_get_should_be_skipped:            true,
			},
		},
		{
			name: "given_valid_slug_with_querier_getArticle_returns_unknown_error_should_return_internal_error",
			args: testData1.args,
			mocks: mocks{
				querier: mock_querier{
					getArticle_args_slug:     testData1.mocks.querier.getArticle_args_slug,
					getArticle_returns_error: dummyErr,
				},
			},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
			configs: configs{
				querier_listArticleTags_should_be_skipped: true,
				userRepo_get_should_be_skipped:            true,
			},
		},
		{
			name: "given_valid_slug_with_querier_listArtileTags_returns_unknown_error_should_return_internal_error",
			args: testData1.args,
			mocks: mocks{
				querier: mock_querier{
					getArticle_args_slug:          testData1.mocks.querier.getArticle_args_slug,
					getArticle_returns_article:    testData1.mocks.querier.getArticle_returns_article,
					listArticleTags_args_slugs:    testData1.mocks.querier.listArticleTags_args_slugs,
					listArticleTags_returns_error: dummyErr,
				},
			},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
			configs: configs{
				userRepo_get_should_be_skipped: true,
			},
		},
		{
			name: "given_valid_slug_with_userRepo_get_returns_error_should_return_not_internal_error",
			args: testData1.args,
			mocks: mocks{
				querier: testData1.mocks.querier,
				userRepository: mock_userRepository{
					get_args_id:       testData1.mocks.userRepository.get_args_id,
					get_returns_error: dummyErr,
				},
			},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			querier := mock_psql_sqlc_gen.NewMockQuerier(t)
			manager := mock_psql_sqlc.NewMockDBManager(t)
			userRepo := mock_auth_repository.NewMockUser(t)
			transaction := mock_transaction.NewMockTransaction(t)
			articleSelector := mock_psql_sqlc.NewMockRawSelector[gen.Article](t)

			if !tt.configs.dbManager_querier_should_be_skipped {
				manager.EXPECT().Querier(mock.Anything).Return(querier)
			}
			if !tt.configs.querier_getArticle_should_be_skipped {
				querier.EXPECT().
					GetArticle(
						mock.Anything,
						tt.mocks.querier.getArticle_args_slug,
					).
					Return(
						tt.mocks.querier.getArticle_returns_article,
						tt.mocks.querier.getArticle_returns_error,
					)
			}
			if !tt.configs.querier_listArticleTags_should_be_skipped {
				querier.EXPECT().
					ListArticleTags(
						mock.Anything,
						tt.mocks.querier.listArticleTags_args_slugs,
					).
					Return(
						tt.mocks.querier.listArticleTags_returns_articleTags,
						tt.mocks.querier.listArticleTags_returns_error,
					)
			}
			if !tt.configs.userRepo_get_should_be_skipped {
				userRepo.EXPECT().
					Get(
						mock.Anything,
						tt.mocks.userRepository.get_args_id,
					).
					Return(
						tt.mocks.userRepository.get_returns_user,
						tt.mocks.userRepository.get_returns_error,
					)
			}

			got, err := NewArtile(manager, transaction, userRepo, articleSelector).Get(context.Background(), tt.args.slug)
			assert.Equal(t, tt.wants.article, got)
			if !assert.ErrorIs(t, err, tt.wants.err) {
				t.Logf("%+v", err)
			}
		})
	}
}

func TestArticle_Save(t *testing.T) {
	type args struct {
		article model.Article
	}
	type mock_querier_upsertArticle struct {
		args_params gen.UpsertArticleParams
		returns_err error
	}
	type mock_querier_deleteArticleTagBySlug struct {
		args_slug   uuid.UUID
		returns_err error
	}
	type mock_querier_createArticleTag struct {
		args_params []gen.CreateArticleTagParams
		returns_int int64
		returns_err error
	}
	type mocks struct {
		querier_upsertArticle          mock_querier_upsertArticle
		querier_deleteArticleTagBySlug mock_querier_deleteArticleTagBySlug
		querier_createArticleTag       mock_querier_createArticleTag
	}
	type configs struct {
		transaction_run_should_be_skipped                bool
		dbManager_querier_should_be_skipped              bool
		querier_upsertArticle_should_be_skipped          bool
		querier_deleteArticleTagBySlug_should_be_skipped bool
		querier_createArticleTag_should_be_skipped       bool
	}

	var (
		slug     = uuid.New()
		now      = premitive.NewJSTTime(time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC))
		dummyErr = errors.New("dummy")
		testDat1 = struct {
			args  args
			mocks mocks
			want  error
		}{
			args: args{
				article: model.Article{
					Slug: model.Slug(slug.String()),
					Tags: []model.TagName{
						"tag1", "tag2",
					},
					Author: model.UserProfile{
						ID:    "author",
						Name:  "name",
						Bio:   "bio",
						Image: "image",
					},
					Contents: model.ArticleContents{
						Title:       "title",
						Description: "desc",
						Body:        "body",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			mocks: mocks{
				querier_upsertArticle: mock_querier_upsertArticle{
					args_params: gen.UpsertArticleParams{
						Slug:        slug,
						Author:      "author",
						Body:        "body",
						Title:       "title",
						Description: "desc",
						CreatedAt: pgtype.Timestamptz{
							Time:  now.Time(),
							Valid: true,
						},
						UpdatedAt: pgtype.Timestamptz{
							Time:  now.Time(),
							Valid: true,
						},
					},
				},
				querier_deleteArticleTagBySlug: mock_querier_deleteArticleTagBySlug{
					args_slug: slug,
				},
				querier_createArticleTag: mock_querier_createArticleTag{
					args_params: []gen.CreateArticleTagParams{
						{
							ArticleSlug: slug,
							TagName:     "tag1",
							CreatedAt: pgtype.Timestamptz{
								Time:  now.Time(),
								Valid: true,
							},
							UpdatedAt: pgtype.Timestamptz{
								Time:  now.Time(),
								Valid: true,
							},
						},
						{
							ArticleSlug: slug,
							TagName:     "tag2",
							CreatedAt: pgtype.Timestamptz{
								Time:  now.Time(),
								Valid: true,
							},
							UpdatedAt: pgtype.Timestamptz{
								Time:  now.Time(),
								Valid: true,
							},
						},
					},
					returns_int: 2,
				},
			},
			want: nil,
		}
	)

	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    error
		configs configs
	}{
		{
			name:  "given_valid_article_should_call_expect_function_and_return_nil",
			args:  testDat1.args,
			mocks: testDat1.mocks,
			want:  testDat1.want,
		},
		{
			name: "given_invalid_slug_article_should_return_internal_error",
			args: args{
				article: model.Article{
					Slug: "invalid_slug",
				},
			},
			want: derrors.Errors.Internal.Err,
			configs: configs{
				transaction_run_should_be_skipped:                true,
				dbManager_querier_should_be_skipped:              true,
				querier_upsertArticle_should_be_skipped:          true,
				querier_deleteArticleTagBySlug_should_be_skipped: true,
				querier_createArticleTag_should_be_skipped:       true,
			},
		},
		{
			name: "given_valid_article_with_querier_upsertArticle_return_error_should_return_internal_error",
			args: testDat1.args,
			mocks: mocks{
				querier_upsertArticle: mock_querier_upsertArticle{
					args_params: testDat1.mocks.querier_upsertArticle.args_params,
					returns_err: dummyErr,
				},
			},
			want: derrors.Errors.Internal.Err,
			configs: configs{
				querier_deleteArticleTagBySlug_should_be_skipped: true,
				querier_createArticleTag_should_be_skipped:       true,
			},
		},
		{
			name: "given_valid_article_with_querier_deleteArticleTags_return_error_should_return_internal_error",
			args: testDat1.args,
			mocks: mocks{
				querier_upsertArticle: testDat1.mocks.querier_upsertArticle,
				querier_deleteArticleTagBySlug: mock_querier_deleteArticleTagBySlug{
					args_slug:   testDat1.mocks.querier_deleteArticleTagBySlug.args_slug,
					returns_err: dummyErr,
				},
			},
			want: derrors.Errors.Internal.Err,
			configs: configs{
				querier_createArticleTag_should_be_skipped: true,
			},
		},
		{
			name: "given_valid_article_with_querier_createArtileTags_return_error_should_return_internal_error",
			args: testDat1.args,
			mocks: mocks{
				querier_upsertArticle:          testDat1.mocks.querier_upsertArticle,
				querier_deleteArticleTagBySlug: testDat1.mocks.querier_deleteArticleTagBySlug,
				querier_createArticleTag: mock_querier_createArticleTag{
					args_params: testDat1.mocks.querier_createArticleTag.args_params,
					returns_err: dummyErr,
				},
			},
			want: derrors.Errors.Internal.Err,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbManager := mock_psql_sqlc.NewMockDBManager(t)
			querier := mock_psql_sqlc_gen.NewMockQuerier(t)
			transaction := mock_transaction.NewMockTransaction(t)
			userRepo := mock_auth_repository.NewMockUser(t)
			artileSelector := mock_psql_sqlc.NewMockRawSelector[gen.Article](t)

			if !tt.configs.transaction_run_should_be_skipped {
				expectationsRunTransaction(t, transaction)
			}
			if !tt.configs.dbManager_querier_should_be_skipped {
				dbManager.EXPECT().
					Querier(mock.Anything).
					Return(querier)
			}
			if !tt.configs.querier_upsertArticle_should_be_skipped {
				querier.EXPECT().
					UpsertArticle(
						mock.Anything,
						tt.mocks.querier_upsertArticle.args_params,
					).
					Return(tt.mocks.querier_upsertArticle.returns_err)
			}
			if !tt.configs.querier_deleteArticleTagBySlug_should_be_skipped {
				querier.EXPECT().
					DeleteArticleTagBySlug(
						mock.Anything,
						tt.mocks.querier_deleteArticleTagBySlug.args_slug,
					).
					Return(tt.mocks.querier_deleteArticleTagBySlug.returns_err)
			}
			if !tt.configs.querier_createArticleTag_should_be_skipped {
				querier.EXPECT().
					CreateArticleTag(
						mock.Anything,
						tt.mocks.querier_createArticleTag.args_params,
					).
					Return(
						tt.mocks.querier_createArticleTag.returns_int,
						tt.mocks.querier_createArticleTag.returns_err,
					)
			}

			err := NewArtile(dbManager, transaction, userRepo, artileSelector).Save(context.Background(), tt.args.article)
			if !assert.ErrorIs(t, err, tt.want) {
				t.Logf("%+v", err)
			}
		})
	}
}

func toPtr[T any](v T) *T {
	return &v
}

func TestArticle_Search(t *testing.T) {
	type args struct {
		param repository.ArticleSearchParam
	}
	type mock_articleSelector struct {
		select_args_query    squirrel.SelectBuilder
		select_returns_rows  []gen.Article
		select_returns_error error
	}
	type mock_querier struct {
		listArticleTags_args_slugs          []uuid.UUID
		listArticleTags_returns_articleTags []gen.ListArticleTagsRow
		listArticleTags_returns_error       error
	}
	type mock_userRepository struct {
		list_args_userIDs  []authmodel.UserID
		list_returns_users []authmodel.User
		list_returns_error error
	}
	type mocks struct {
		articleSelector mock_articleSelector
		querier         mock_querier
		userRepository  mock_userRepository
	}
	type wants struct {
		articles model.ArticleSlice
		err      error
	}
	type configs struct {
		dbManager_querier_should_be_skipped       bool
		querier_listArticleTags_should_be_skipped bool
		userRepositry_list_should_be_skipped      bool
	}

	var (
		now       = time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC)
		nowTz     = ToTimestamptz(now)
		nowJST    = premitive.NewJSTTime(now)
		slug1     = uuid.New()
		slug2     = uuid.New()
		dummyErr  = errors.New("dummy")
		testData1 = struct {
			args  args
			mocks mocks
			wants wants
		}{
			args: args{
				param: repository.ArticleSearchParam{
					Tag:         toPtr(model.TagName("tag1")),
					Author:      toPtr(authmodel.UserID("author1")),
					FavoritedBy: toPtr(authmodel.UserID("user1")),
					Limit:       100,
					Offset:      20,
				},
			},
			mocks: mocks{
				articleSelector: mock_articleSelector{
					select_args_query: squirrel.Select("a.*").
						From("article as a").
						LeftJoin("article_favorite as f on a.slug = f.article_slug").
						LeftJoin("article_tag as t on a.slug = t.article_slug").
						Where(squirrel.Eq{
							"a.author":   "author1",
							"f.user_id":  "user1",
							"t.tag_name": "tag1",
						}).
						Limit(100).
						Offset(20),
					select_returns_rows: []gen.Article{
						{
							Slug:        slug1,
							Author:      "author1",
							Body:        "body1",
							Title:       "title1",
							Description: "desc1",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
						{
							Slug:        slug2,
							Author:      "author2",
							Body:        "body2",
							Title:       "title2",
							Description: "desc2",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
					},
				},
				querier: mock_querier{
					listArticleTags_args_slugs: []uuid.UUID{
						slug1, slug2,
					},
					listArticleTags_returns_articleTags: []gen.ListArticleTagsRow{
						{
							ArticleSlug: slug1,
							Name:        "tag1",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
						{
							ArticleSlug: slug1,
							Name:        "tag2",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
						{
							ArticleSlug: slug2,
							Name:        "tag2",
							CreatedAt:   nowTz,
							UpdatedAt:   nowTz,
						},
					},
				},
				userRepository: mock_userRepository{
					list_args_userIDs: []authmodel.UserID{
						"author1", "author2",
					},
					list_returns_users: []authmodel.User{
						{
							ID: "author1",
							Profile: authmodel.Profile{
								Username: "name1",
								Bio:      "bio1",
								Image:    "https://sample.com/image1.png",
							},
						},
						{
							ID: "author2",
							Profile: authmodel.Profile{
								Username: "name2",
								Bio:      "bio2",
								Image:    "https://sample.com/image2.png",
							},
						},
					},
				},
			},
			wants: wants{
				articles: model.ArticleSlice{
					{
						Slug: model.Slug(slug1.String()),
						Tags: []model.TagName{
							"tag1", "tag2",
						},
						Author: model.UserProfile{
							ID:    "author1",
							Name:  "name1",
							Bio:   "bio1",
							Image: "https://sample.com/image1.png",
						},
						Contents: model.ArticleContents{
							Title:       "title1",
							Description: "desc1",
							Body:        "body1",
						},
						CreatedAt: nowJST,
						UpdatedAt: nowJST,
					},
					{
						Slug: model.Slug(slug2.String()),
						Tags: []model.TagName{
							"tag2",
						},
						Author: model.UserProfile{
							ID:    "author2",
							Name:  "name2",
							Bio:   "bio2",
							Image: "https://sample.com/image2.png",
						},
						Contents: model.ArticleContents{
							Title:       "title2",
							Description: "desc2",
							Body:        "body2",
						},
						CreatedAt: nowJST,
						UpdatedAt: nowJST,
					},
				},
			},
		}
	)

	tests := []struct {
		name    string
		args    args
		mocks   mocks
		wants   wants
		configs configs
	}{
		{
			name:  "valid_params_should_call_expect_functions_and_returns_expect_articles",
			args:  testData1.args,
			mocks: testData1.mocks,
			wants: testData1.wants,
		},
		{
			name: "valid_params_with_articleSelector_select_returns_error_should_returns_internal_error",
			args: testData1.args,
			mocks: mocks{
				articleSelector: mock_articleSelector{
					select_args_query:    testData1.mocks.articleSelector.select_args_query,
					select_returns_error: dummyErr,
				},
			},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
			configs: configs{
				dbManager_querier_should_be_skipped:       true,
				querier_listArticleTags_should_be_skipped: true,
				userRepositry_list_should_be_skipped:      true,
			},
		},
		{
			name: "valid_params_with_querier_listArticleTags_returns_error_should_returns_internal_error",
			args: testData1.args,
			mocks: mocks{
				articleSelector: testData1.mocks.articleSelector,
				querier: mock_querier{
					listArticleTags_args_slugs:    testData1.mocks.querier.listArticleTags_args_slugs,
					listArticleTags_returns_error: dummyErr,
				},
			},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
			configs: configs{
				userRepositry_list_should_be_skipped: true,
			},
		},
		{
			name: "valid_params_with_userRepo_list_returns_error_should_returns_internal_error",
			args: testData1.args,
			mocks: mocks{
				articleSelector: testData1.mocks.articleSelector,
				querier:         testData1.mocks.querier,
				userRepository: mock_userRepository{
					list_args_userIDs:  testData1.mocks.userRepository.list_args_userIDs,
					list_returns_error: dummyErr,
				},
			},
			wants: wants{
				err: derrors.Errors.Internal.Err,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbManager := mock_psql_sqlc.NewMockDBManager(t)
			executor := mock_psql_sqlc.NewMockContextExecutor(t)
			transaction := mock_transaction.NewMockTransaction(t)
			querier := mock_psql_sqlc_gen.NewMockQuerier(t)
			userRepo := mock_auth_repository.NewMockUser(t)
			articleSelector := mock_psql_sqlc.NewMockRawSelector[gen.Article](t)

			dbManager.EXPECT().Executor(mock.Anything).Return(executor)

			articleSelector.EXPECT().
				Select(
					mock.Anything,
					executor,
					tt.mocks.articleSelector.select_args_query,
				).
				Return(
					tt.mocks.articleSelector.select_returns_rows,
					tt.mocks.articleSelector.select_returns_error,
				)
			if !tt.configs.dbManager_querier_should_be_skipped {
				dbManager.EXPECT().Querier(mock.Anything).Return(querier)
			}
			if !tt.configs.querier_listArticleTags_should_be_skipped {
				querier.EXPECT().
					ListArticleTags(
						mock.Anything,
						tt.mocks.querier.listArticleTags_args_slugs,
					).
					Return(
						tt.mocks.querier.listArticleTags_returns_articleTags,
						tt.mocks.querier.listArticleTags_returns_error,
					)
			}
			if !tt.configs.userRepositry_list_should_be_skipped {
				userRepo.EXPECT().
					List(
						mock.Anything,
						lo.ToAnySlice(tt.mocks.userRepository.list_args_userIDs)...,
					).
					Return(
						tt.mocks.userRepository.list_returns_users,
						tt.mocks.userRepository.list_returns_error,
					)
			}

			got, err := NewArtile(dbManager, transaction, userRepo, articleSelector).Search(context.Background(), tt.args.param)
			assert.Equal(t, tt.wants.articles, got)
			if !assert.ErrorIs(t, err, tt.wants.err) {
				t.Logf("%+v", err)
			}
		})
	}
}
