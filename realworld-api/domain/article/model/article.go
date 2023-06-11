package model

import (
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
)

var (
	newValidateOnce sync.Once
	validate        *validator.Validate
)

func getValidate() *validator.Validate {
	newValidateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}

type Slug premitive.UID

func NewSlug(s string) (Slug, error) {
	slug, err := premitive.NewUID(s)
	if err != nil {
		return "", err
	}
	return Slug(slug), nil
}

func (s Slug) String() string {
	return premitive.UID(s).String()
}

type Article struct {
	Slug      Slug             `validate:"required"`
	Author    authmodel.UserID `validate:"required"`
	Contents  ArticleContents
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func NewArticle(slug Slug, contents ArticleContents, author authmodel.UserID) (*Article, error) {
	return &Article{
		Slug:      slug,
		Contents:  contents,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (a *Article) SetContents(contents ArticleContents) {
	a.Contents = contents
	a.UpdatedAt = time.Now()
}

type ArticleContents struct {
	Title       premitive.Title `validate:"required"`
	Description premitive.LongText
	Body        premitive.LongText
}

func NewArticleContents(title premitive.Title, description, body premitive.LongText) (*ArticleContents, error) {
	contents := &ArticleContents{
		Title:       title,
		Description: description,
		Body:        body,
	}

	// TODO
	if err := getValidate().Struct(contents); err != nil {
		return nil, err
	}
	return contents, nil
}
