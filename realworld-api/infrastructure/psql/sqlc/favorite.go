package sqlc

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
)

type Favorite struct{}

var _ repository.Favorite = (*Favorite)(nil)

func NewFavorite() *Favorite {
	return &Favorite{}
}

func (f *Favorite) ListBySlug(_ context.Context, _ model.Slug) (model.FavoriteSlice, error) {
	panic("not implemented") // TODO: Implement
}

func (f *Favorite) ListBySlugs(_ context.Context, _ ...model.Slug) (model.FavoriteSliceMap, error) {
	return model.FavoriteSliceMap{}, nil
}
