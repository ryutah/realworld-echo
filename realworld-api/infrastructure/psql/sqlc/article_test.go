package sqlc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	. "github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc/gen"
	mock_psql_sqlc "github.com/ryutah/realworld-echo/realworld-api/internal/mock/psql/sqlc"
	mock_psql_sqlc_gen "github.com/ryutah/realworld-echo/realworld-api/internal/mock/psql/sqlc/gen"
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
			got, err := NewArtile(nil).GenerateID(context.Background())
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
		listArticleTags_args_slugs          []string
		listArticleTags_returns_articleTags []gen.ListArticleTagsRow
		listArticleTags_returns_error       error
	}
	type mocks struct {
		querier mock_querier
	}
	type wants struct {
		article *model.Article
		err     error
	}
	type configs struct {
		dbManager_querier_should_be_skipped       bool
		querier_getArticle_should_be_skipped      bool
		querier_listArticleTags_should_be_skipped bool
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
					listArticleTags_args_slugs: []string{
						slug.String(),
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
			},
			wants: wants{
				article: &model.Article{
					Slug: model.Slug(slug.String()),
					Tags: []model.TagName{
						"tag1", "tag2",
					},
					Author: "author",
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
			},
		},
		{
			name: "given_valid_slug_with_querier_listArtileTags_returns_unknown_error_should_return_not_found_error",
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			querier := mock_psql_sqlc_gen.NewMockQuerier(t)
			manager := mock_psql_sqlc.NewMockDBManager(t)

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

			got, err := NewArtile(manager).Get(context.Background(), tt.args.slug)
			assert.Equal(t, tt.wants.article, got)
			if !assert.ErrorIs(t, err, tt.wants.err) {
				t.Logf("%+v", err)
			}
		})
	}
}
