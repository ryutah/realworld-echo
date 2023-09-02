package model

import (
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
)

type Follow struct {
	User      UserProfile
	Follower  UserProfile
	CreatedAt premitive.JSTTime
	UpdateAt  premitive.JSTTime
}

type FollowersExistsMap map[authmodel.UserID]bool

func (f FollowersExistsMap) IsFollowing(follower authmodel.UserID) bool {
	return f[follower]
}
