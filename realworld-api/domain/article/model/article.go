package model

import (
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xvalidator"
	"github.com/samber/lo"
)

type ArticleSlice []Article

func (a ArticleSlice) Slugs() []Slug {
	return lo.Map(a, func(item Article, _ int) Slug {
		return item.Slug
	})
}

type Article struct {
	Slug      Slug             `validate:"required"`
	Author    authmodel.UserID `validate:"required"`
	Contents  ArticleContents
	CreatedAt premitive.JSTTime `validate:"required"`
	UpdatedAt premitive.JSTTime `validate:"required"`
}

func NewArticle(slug Slug, contents ArticleContents, author authmodel.UserID) (*Article, error) {
	article := &Article{
		Slug:      slug,
		Contents:  contents,
		Author:    author,
		CreatedAt: premitive.NewJSTTime(xtime.Now()),
		UpdatedAt: premitive.NewJSTTime(xtime.Now()),
	}
	if err := xvalidator.Validator().Struct(article); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return article, nil
}

func (a *Article) Edit(contents ArticleContents) {
	a.Contents = contents
	a.UpdatedAt = premitive.NewJSTTime(xtime.Now())
}

type (
	Slug premitive.UID
)

func NewSlug(s string) (Slug, error) {
	uid, err := premitive.NewUID(s)
	if err != nil {
		return "", err
	}

	return Slug(uid), nil
}

func (s Slug) String() string {
	return premitive.UID(s).String()
}

type ArticleContents struct {
	Tags        []ArticleTag
	Title       premitive.Title `validate:"required"`
	Description premitive.LongText
	Body        premitive.LongText
}

func NewArticleContents(title, description, body string, tags []ArticleTag) (*ArticleContents, error) {
	ptitle, err := premitive.NewTitle(title)
	if err != nil {
		return nil, err
	}
	pdescription, err := premitive.NewLongText(description)
	if err != nil {
		return nil, err
	}
	pbody, err := premitive.NewLongText(body)
	if err != nil {
		return nil, err
	}

	contents := &ArticleContents{
		Tags:        tags,
		Title:       ptitle,
		Description: pdescription,
		Body:        pbody,
	}

	if err := xvalidator.Validator().Struct(contents); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return contents, nil
}

type ArticleTag struct {
	Tag premitive.ShortText `validate:"required"`
}

func NewArticleTag(tag string) (*ArticleTag, error) {
	ptag, err := premitive.NewShortText(tag)
	if err != nil {
		return nil, err
	}

	articleTag := &ArticleTag{
		Tag: ptag,
	}
	if err := xvalidator.Validator().Struct(articleTag); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return articleTag, nil
}
