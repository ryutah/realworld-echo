package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type ArticleSearchParam struct {
	Tag         *model.ArticleTag
	Author      *authmodel.UserID
	FavoritedBy *authmodel.UserID
}

type Article interface {
	GenerateID(context.Context) (model.Slug, error)
	Get(context.Context, model.Slug) (*model.Article, error)
	Save(context.Context, model.Article) error
	Search(context.Context, ArticleSearchParam) ([]model.Article, error)
}
