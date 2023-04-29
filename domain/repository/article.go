package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/domain/model"
	"github.com/ryutah/realworld-echo/domain/premitive"
)

type Article interface {
	GenerateID(context.Context) (model.ArticleID, error)
	Get(context.Context, premitive.Slug) (*model.Article, error)
	Create(context.Context, model.Article) error
}
