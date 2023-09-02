package sqlc

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
)

type Favorite struct{}

var _ repository.Favorite = (*Favorite)(nil)

func NewFavorite() *Favorite {
	return &Favorite{}
}

func (f *Favorite) ListBySlug(_ context.Context, slug model.Slug) (model.FavoriteSlice, error) {
	return model.FavoriteSlice{
		model.Favorite{
			ArticleSlug: slug,
			UserID:      "user1",
			CreatedAt:   premitive.NewJSTTime(xtime.Now()),
			UpdatedAt:   premitive.NewJSTTime(xtime.Now()),
		},
	}, nil
}

func (f *Favorite) ListBySlugs(_ context.Context, _ ...model.Slug) (model.FavoriteSliceMap, error) {
	return model.FavoriteSliceMap{}, nil
}
