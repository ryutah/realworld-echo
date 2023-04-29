package model

import (
	"time"

	"github.com/ryutah/realworld-echo/domain/premitive"
)

type ArticleID string

type Article struct {
	ID          ArticleID
	Slug        premitive.Slug
	Title       premitive.Title
	Description premitive.LongText
	Body        premitive.LongText
	Author      *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewArticle(
	id ArticleID,
	slug premitive.Slug,
	title premitive.Title,
	desciption premitive.LongText,
	author *User,
	body premitive.LongText,
) (*Article, error) {
	return &Article{
		ID:          id,
		Slug:        slug,
		Description: desciption,
		Body:        body,
		Author:      author,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
