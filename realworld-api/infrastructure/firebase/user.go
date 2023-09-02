package firebase

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/repository"
)

type User struct{}

var _ repository.User = (*User)(nil)

func NewUser() *User {
	return &User{}
}

func (u *User) Get(_ context.Context, _ model.UserID) (*model.User, error) {
	return &model.User{
		ID: "user1",
		Account: model.Account{
			Email: "example.com@gmail.com",
		},
		Profile: model.Profile{
			Username: "name",
			Bio:      "bio",
			Image:    "http://example.com/image.png",
		},
	}, nil
}

func (u *User) List(_ context.Context, _ ...model.UserID) ([]model.User, error) {
	panic("not implemented") // TODO: Implement
}
