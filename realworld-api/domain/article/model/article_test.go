package model_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/stretchr/testify/assert"
)

func TestNewArticleContents(t *testing.T) {
	type args struct {
		title       string
		description string
		body        string
	}
	type wants struct {
		contents *ArticleContents
		err      error
	}

	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "valid_arguments_should_return_expected_contents",
			args: args{
				title:       "title",
				description: "description",
				body:        "body",
			},
			want: wants{
				contents: &ArticleContents{
					Title:       "title",
					Description: "description",
					Body:        "body",
				},
				err: nil,
			},
		},
		{
			name: "rquired_field_only_should_return_expected_contents",
			args: args{
				title:       "title",
				description: "",
				body:        "",
			},
			want: wants{
				contents: &ArticleContents{
					Title:       "title",
					Description: "",
					Body:        "",
				},
				err: nil,
			},
		},
		{
			name: "blank_title_should_return_validation_error",
			args: args{
				title: "",
			},
			want: wants{
				contents: nil,
				err:      errors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_title_should_return_validation_error",
			args: args{
				title: strings.Repeat("a", 10000),
			},
			want: wants{
				contents: nil,
				err:      errors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_description_should_return_validation_error",
			args: args{
				title:       "title",
				description: strings.Repeat("a", 10000),
			},
			want: wants{
				contents: nil,
				err:      errors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_body_should_return_validation_error",
			args: args{
				title: "title",
				body:  strings.Repeat("a", 10000),
			},
			want: wants{
				contents: nil,
				err:      errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewArticleContents(tt.args.title, tt.args.description, tt.args.body)
			assert.Equal(t, tt.want.contents, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}

func TestNewArticle(t *testing.T) {
	type args struct {
		slug     Slug
		contents ArticleContents
		author   UserProfile
		tags     []TagName
	}
	type mocks struct {
		now func() time.Time
	}
	type wants struct {
		article *Article
		err     error
	}

	var (
		contents = ArticleContents{
			Title:       "title",
			Description: "description",
			Body:        "body",
		}
		tags    = []TagName{"tag1"}
		now     = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		nowFunc = func() time.Time { return now }
	)

	tests := []struct {
		name  string
		args  args
		mocks mocks
		want  wants
	}{
		{
			name: "valid_arguments_should_return_expected_article",
			args: args{
				slug:     "slug",
				contents: contents,
				author: UserProfile{
					ID: "author",
				},
				tags: tags,
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				article: &Article{
					Slug: "slug",
					Author: UserProfile{
						ID: "author",
					},
					Contents:  contents,
					Tags:      tags,
					CreatedAt: premitive.NewJSTTime(now),
					UpdatedAt: premitive.NewJSTTime(now),
				},
				err: nil,
			},
		},
		{
			name: "blank_slug_should_return_validation_error",
			args: args{
				slug:     "",
				contents: contents,
				author: UserProfile{
					ID: "author",
				},
				tags: tags,
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				article: nil,
				err:     errors.Errors.Validation.Err,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset := xtime.SetNowFunc(tt.mocks.now)
			defer reset()

			got, err := NewArticle(tt.args.slug, tt.args.contents, tt.args.author, tt.args.tags)
			assert.Equal(t, tt.want.article, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}

func TestNewSlug(t *testing.T) {
	type args struct {
		s string
	}
	type wants struct {
		slug string
		err  error
	}

	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "valid_uid_should_return_expected_slug",
			args: args{
				s: "slug",
			},
			want: wants{
				slug: "slug",
			},
		},
		{
			name: "invalid_uid_should_return_validation_error",
			args: args{
				s: strings.Repeat("a", 10000),
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
		{
			name: "blank_uid_should_return_validation_error",
			args: args{
				s: "",
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSlug(tt.args.s)
			assert.Equal(t, tt.want.slug, got.String())
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}

func TestArticle_Edit(t *testing.T) {
	type args struct {
		contents ArticleContents
	}
	type mocks struct {
		now func() time.Time
	}
	type wants struct {
		article *Article
	}

	var (
		beforeNow   = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
		afterNow    = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		nowFunc     = func() time.Time { return afterNow }
		baseArticle = func() *Article {
			return &Article{
				Slug:   "slug",
				Author: UserProfile{ID: "author"},
				Contents: ArticleContents{
					Title:       "title",
					Description: "description",
					Body:        "body",
				},
				CreatedAt: premitive.JSTTime(beforeNow),
				UpdatedAt: premitive.JSTTime(beforeNow),
			}
		}
		updatedArticle = func(article *Article, newContents ArticleContents, newUpdatedAt premitive.JSTTime) *Article {
			article.Contents = newContents
			article.UpdatedAt = newUpdatedAt
			return article
		}
	)

	tests := []struct {
		name   string
		target *Article
		args   args
		mocks  mocks
		want   wants
	}{
		{
			name:   "valid_contents_should_update_article_content_and_updated_at",
			target: baseArticle(),
			args: args{
				contents: ArticleContents{
					Title:       "new_title",
					Description: "new_desc",
					Body:        "new_bofy",
				},
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				article: updatedArticle(
					baseArticle(),
					ArticleContents{
						Title:       "new_title",
						Description: "new_desc",
						Body:        "new_bofy",
					},
					premitive.NewJSTTime(afterNow),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset := xtime.SetNowFunc(tt.mocks.now)
			defer reset()

			tt.target.Edit(tt.args.contents)
			assert.Equal(t, tt.want.article, tt.target)
		})
	}
}

func TestArticleSlice_Slugs(t *testing.T) {
	tests := []struct {
		name string
		a    ArticleSlice
		want []Slug
	}{
		{
			name: "should_return_expected_slugs",
			a: ArticleSlice{
				{Slug: "slug1"},
				{Slug: "slug2"},
				{Slug: "slug3"},
			},
			want: []Slug{"slug1", "slug2", "slug3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Slugs()
			assert.Equal(t, tt.want, got)
		})
	}
}
