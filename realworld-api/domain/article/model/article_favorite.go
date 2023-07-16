package model

import (
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xvalidator"
)

type ArticleFavorite struct {
	ArticleSlug Slug              `validate:"required"`
	UserID      authmodel.UserID  `validate:"required"`
	CreatedAt   premitive.JSTTime `validate:"required"`
	UpdateAt    premitive.JSTTime `validate:"required"`
}

func NewArticleFavorite(slug Slug, userID authmodel.UserID) (*ArticleFavorite, error) {
	fav := ArticleFavorite{
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
