package model

import (
	"time"

	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
)

type CommentID premitive.UID

type Comment struct {
	ID        CommentID          `validate:"required"`
	Slug      Slug               `validate:"required"`
	Author    authmodel.UserID   `validate:"required"`
	Body      premitive.LongText `validate:"required"`
	CreatedAt time.Time          `validate:"required"`
	UpdatedAt time.Time          `validate:"required"`
}
