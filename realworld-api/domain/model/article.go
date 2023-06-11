package model

import (
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
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

type ArticleID premitive.UID

type Article struct {
	ID        ArticleID
	Contents  ArticleContents
	Author    *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(id ArticleID, contents ArticleContents, author *User) (*Article, error) {
	return &Article{
		ID:        id,
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
	Slug        premitive.Slug  `validate:"required"`
	Title       premitive.Title `validate:"required"`
	Description premitive.LongText
	Body        premitive.LongText
}

func NewArticleContents(slug, title, description, body string) (*ArticleContents, error) {
	pslag, err := premitive.NewSlug(slug)
	if err != nil {
		return nil, err
	}
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
		Slug:        pslag,
		Title:       ptitle,
		Description: pdescription,
		Body:        pbody,
	}

	if err := getValidate().Struct(contents); err != nil {
		return nil, nil
	}
	return contents, nil
}
