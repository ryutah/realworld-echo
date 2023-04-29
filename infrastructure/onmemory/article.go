package onmemory

import (
	"context"
	"time"

	"github.com/ryutah/realworld-echo/domain/model"
	"github.com/ryutah/realworld-echo/domain/premitive"
	"github.com/ryutah/realworld-echo/domain/repository"
)

type Article struct{}

func NewArticle() repository.Article {
	return &Article{}
}

func (a *Article) GenerateID(ctx context.Context) (model.ArticleID, error) {
	return "", nil
}

func (a *Article) Get(_ context.Context, slug premitive.Slug) (*model.Article, error) {
	return &model.Article{
		Slug:        slug,
		Title:       "dummytitle",
		Description: "dummyDescription",
		Body:        "dummyBody",
		Author: &model.User{
			Account: &model.Account{
				Email: "aaa@gmail.com",
			},
			Profile: &model.Profile{
				Username: "sample",
				Image:    "http:/xxx.com",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (a *Article) Create(_ context.Context, _ model.Article) error {
	panic("not implemented") // TODO: Implement
}
