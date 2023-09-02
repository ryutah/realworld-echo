package model

import (
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
)

type UserProfile struct {
	ID    authmodel.UserID `validate:"required"`
	Name  premitive.Name
	Bio   premitive.ShortText
	Image premitive.URL
}

func NewUserProfile(u authmodel.User) *UserProfile {
	return &UserProfile{
		ID:    u.ID,
		Name:  u.Profile.Username,
		Bio:   u.Profile.Bio,
		Image: u.Profile.Image,
	}
}

func ReCreateUserProfile(idStr, nameStr, bioStr, imageStr string) (*UserProfile, error) {
	id, err := authmodel.NewUserID(idStr)
	if err != nil {
		return nil, err
	}
	name, err := premitive.NewName(nameStr)
	if err != nil {
		return nil, err
	}
	bio, err := premitive.NewShortText(bioStr)
	if err != nil {
		return nil, err
	}
	image, err := premitive.NewURL(imageStr)
	if err != nil {
		return nil, err
	}

	return &UserProfile{
		ID:    id,
		Name:  name,
		Bio:   bio,
		Image: image,
	}, nil
}
