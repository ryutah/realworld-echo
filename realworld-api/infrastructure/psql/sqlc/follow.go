package sqlc

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/samber/lo"
)

type Follow struct {
	repository.Follow
}

var _ repository.Follow = (*Follow)(nil)

func NewFollow() *Follow {
	return &Follow{}
}

func (f *Follow) ExistsList(ctx context.Context, followedBy authmodel.UserID, following ...authmodel.UserID) (model.FollowersExistsMap, error) {
	maps := make(model.FollowersExistsMap)
	lo.ForEach(following, func(item authmodel.UserID, _ int) {
		maps[item] = false
	})
	return maps, nil
}
