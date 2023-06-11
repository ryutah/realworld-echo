package model

import (
	"time"

	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type Favorite struct {
	UserID    authmodel.UserID `validate:"required"`
	Slug      Slug             `validate:"required"`
	CreatedAt time.Time        `validate:"required"`
	UpdatedAt time.Time        `validate:"required"`
}

func NewFavorite(userID authmodel.UserID, slug Slug) *Favorite {
	return &Favorite{
		UserID:    userID,
		Slug:      slug,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
