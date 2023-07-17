package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
)

type Favorite interface {
	ListBySlug(context.Context, model.Slug) (model.FavoriteSlice, error)
}
