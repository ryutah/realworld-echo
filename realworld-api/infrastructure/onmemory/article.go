package onmemory

import (
	"context"
	"time"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type Article struct {
	repository.Article
}

func NewArticle() repository.Article {
	return &Article{}
}

func (a *Article) GenerateID(ctx context.Context) (model.Slug, error) {
	return "", nil
}

func (a *Article) Get(_ context.Context, slug model.Slug) (*model.Article, error) {
	return &model.Article{
		Slug: slug,
		Contents: model.ArticleContents{
			Title:       "dummytitle",
			Description: "dummyDescription",
			Body:        "dummyBody",
		},
		Author:    authmodel.UserID("dummy"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (a *Article) Create(_ context.Context, _ model.Article) error {
	panic("not implemented") // TODO: Implement
}
