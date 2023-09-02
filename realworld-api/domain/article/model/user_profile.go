package model

import (
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
)

type UserProfile struct {
	ID    authmodel.UserID
	Name  premitive.Name
	Bio   premitive.ShortText
	Image premitive.URL
}

func NewUserProfile(u authmodel.User) (*UserProfile, error) {
	return &UserProfile{
		ID:    u.ID,
		Name:  u.Profile.Username,
		Bio:   u.Profile.Bio,
		Image: u.Profile.Image,
	}, nil
}
