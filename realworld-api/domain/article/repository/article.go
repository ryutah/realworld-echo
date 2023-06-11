package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
)

type Article interface {
	GenerateID(context.Context) (model.Slug, error)
	Get(context.Context, model.Slug) (*model.Article, error)
	Save(context.Context, model.Article) error
}
