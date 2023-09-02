package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type Follow interface {
	ExistsList(ctx context.Context, followedBy authmodel.UserID, following ...authmodel.UserID) (model.FollowersExistsMap, error)
}
