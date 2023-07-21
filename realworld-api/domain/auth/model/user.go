package model

import (
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xvalidator"
)

type UserID premitive.UID

func NewUserID(s string) (UserID, error) {
	uid, err := premitive.NewUID(s)
	if err != nil {
		return "", err
	}
	return UserID(uid), nil
}

func (u UserID) String() string {
	return premitive.UID(u).String()
}

type User struct {
	ID      UserID `validate:"required"`
	Account Account
	Profile Profile
}

func NewUser(id UserID, account Account, profile Profile) (*User, error) {
	u := &User{
		ID:      id,
		Account: account,
		Profile: profile,
	}
	if err := xvalidator.Validator().Struct(u); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return u, nil
}

type Account struct {
	Email premitive.Email `validate:"required"`
}

func NewAccount(email string) (*Account, error) {
	e, err := premitive.NewEmail(email)
	if err != nil {
		return nil, err
	}

	account := Account{
		Email: e,
	}
	if err := xvalidator.Validator().Struct(account); err != nil {
		return nil, errors.NewValidationError(0, err)
	}

	return &account, nil
}

type Profile struct {
	Username premitive.Name
	Bio      premitive.ShortText
	Image    premitive.URL
}

func NewProfile(username, image string) (*Profile, error) {
	n, err := premitive.NewName(username)
	if err != nil {
		return nil, err
	}
	u, err := premitive.NewURL(image)
	if err != nil {
		return nil, err
	}

	return &Profile{
		Username: n,
		Image:    u,
	}, nil
}
