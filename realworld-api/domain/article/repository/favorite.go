package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type Favorite interface {
	Count(context.Context, model.Slug) (int, error)
	CountList(context.Context, ...model.Slug) (model.FavoriteCountMap, error)
	Exists(context.Context, authmodel.UserID, model.Slug) (bool, error)
	ExistsList(context.Context, authmodel.UserID, ...model.Slug) (model.FavoriteExistsMap, error)
}
