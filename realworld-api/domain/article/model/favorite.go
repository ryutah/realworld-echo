package model

import (
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xvalidator"
	"github.com/samber/lo"
)

type Favorite struct {
	ArticleSlug Slug              `validate:"required"`
	UserID      authmodel.UserID  `validate:"required"`
	CreatedAt   premitive.JSTTime `validate:"required"`
	UpdateAt    premitive.JSTTime `validate:"required"`
}

func NewFavorite(slug Slug, userID authmodel.UserID) (*Favorite, error) {
	fav := Favorite{
		ArticleSlug: slug,
		UserID:      userID,
		CreatedAt:   premitive.NewJSTTime(xtime.Now()),
		UpdateAt:    premitive.NewJSTTime(xtime.Now()),
	}
	if err := xvalidator.Validator().Struct(fav); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return &fav, nil
}

type FavoriteSlice []Favorite

func (f FavoriteSlice) IsFavorited(userID authmodel.UserID, slug Slug) bool {
	return lo.ContainsBy(f, func(item Favorite) bool {
		return item.UserID == userID && item.ArticleSlug == slug
	})
}
