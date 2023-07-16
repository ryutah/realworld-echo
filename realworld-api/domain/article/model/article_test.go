package model_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/stretchr/testify/assert"
)

func TestNewArticleTag(t *testing.T) {
	type args struct {
		tag string
	}
	type wants struct {
		tag *ArticleTag
		err error
	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "valid_tag_should_return_expecte_tag",
			args: args{
				tag: "valid_tag",
			},
			want: wants{
				tag: &ArticleTag{Tag: "valid_tag"},
				err: nil,
			},
		},
		{
			name: "invalid_tag_should_return_validation_error",
			args: args{
				tag: strings.Repeat("a", 5000),
			},
			want: wants{
				tag: nil,
				err: errors.Errors.Validation.Err,
			},
		},
		{
			name: "blank_tag_should_return_validation_error",
			args: args{
				tag: "",
			},
			want: wants{
				tag: nil,
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewArticleTag(tt.args.tag)

			assert.Equal(t, tt.want.tag, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}

func TestNewArticleContents(t *testing.T) {
	type args struct {
		title       string
		description string
		body        string
		tags        []ArticleTag
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
				tags: []ArticleTag{
					{Tag: "tag1"},
					{Tag: "tag2"},
				},
			},
			want: wants{
				contents: &ArticleContents{
					Title:       "title",
					Description: "description",
					Body:        "body",
					Tags: []ArticleTag{
						{Tag: "tag1"},
						{Tag: "tag2"},
					},
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
				tags:        nil,
			},
			want: wants{
				contents: &ArticleContents{
					Title:       "title",
					Description: "",
					Body:        "",
					Tags:        nil,
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
			got, err := NewArticleContents(tt.args.title, tt.args.description, tt.args.body, tt.args.tags)
			assert.Equal(t, tt.want.contents, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}

func TestNewArticle(t *testing.T) {
	type args struct {
		slug     Slug
		contents ArticleContents
		author   authmodel.UserID
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
			Tags: []ArticleTag{
				{Tag: "tag1"},
			},
			Title:       "title",
			Description: "description",
			Body:        "body",
		}
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
				author:   "author",
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				article: &Article{
					Slug:      "slug",
					Author:    "author",
					Contents:  contents,
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
				author:   "author",
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				article: nil,
				err:     errors.Errors.Validation.Err,
			},
		},
		{
			name: "blank_author_should_return_validation_error",
			args: args{
				slug:     "slug",
				contents: contents,
				author:   "",
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

			got, err := NewArticle(tt.args.slug, tt.args.contents, tt.args.author)
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
			name: "invalid_uid_should_return_expected_slug",
			args: args{
				s: strings.Repeat("a", 10000),
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
				Author: "author",
				Contents: ArticleContents{
					Tags: []ArticleTag{
						{Tag: "tag1"},
						{Tag: "tag2"},
					},
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
					Tags: []ArticleTag{
						{Tag: "new_tag1"},
						{Tag: "new_tag2"},
					},
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
						Tags: []ArticleTag{
							{Tag: "new_tag1"},
							{Tag: "new_tag2"},
						},
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
