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

func (a ArticleSlice) Authors() []authmodel.UserID {
	return lo.Map(a, func(item Article, _ int) authmodel.UserID {
		return item.Author.ID
	})
}

type Article struct {
	Slug      Slug `validate:"required"`
	Tags      []TagName
	Author    UserProfile
	Contents  ArticleContents
	CreatedAt premitive.JSTTime `validate:"required"`
	UpdatedAt premitive.JSTTime `validate:"required"`
}

func NewArticle(slug Slug, contents ArticleContents, author UserProfile, tags []TagName) (*Article, error) {
	now := premitive.NewJSTTime(xtime.Now())
	return ReCreateArticle(slug, contents, author, tags, now, now)
}

func ReCreateArticle(
	slug Slug,
	contents ArticleContents,
	author UserProfile,
	tags []TagName,
	createdAt, updatedAt premitive.JSTTime,
) (*Article, error) {
	article := &Article{
		Slug:      slug,
		Contents:  contents,
		Author:    author,
		Tags:      tags,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
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
	slug := Slug(uid)
	if err := xvalidator.Validator().Var(slug, "required"); err != nil {
		return "", errors.NewValidationError(0, err)
	}
	return slug, nil
}

func (s Slug) String() string {
	return premitive.UID(s).String()
}

type ArticleContents struct {
	Title       premitive.Title `validate:"required"`
	Description premitive.LongText
	Body        premitive.LongText
}

func NewArticleContents(title, description, body string) (*ArticleContents, error) {
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
		Title:       ptitle,
		Description: pdescription,
		Body:        pbody,
	}

	if err := xvalidator.Validator().Struct(contents); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return contents, nil
}
