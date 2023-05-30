package service

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/model"
)

type Auth interface {
	CurrentUserFromToken(context.Context, string) (*model.User, error)
	CurrentUser(context.Context) (user *model.User, ok bool, err error)
}

type auth struct{}

func NewAuth() Auth {
	return &auth{}
}

func (a *auth) CurrentUserFromToken(ctx context.Context, token string) (*model.User, error) {
	return &model.User{
		Account: &model.Account{
			Email: "hogehoge@sample.com",
		},
		Profile: &model.Profile{
			Username: "sample_user",
			Image:    "http://xxxxxxxx.com",
		},
	}, nil
}

func (a *auth) CurrentUser(context.Context) (*model.User, bool, error) {
	return &model.User{
		Account: &model.Account{
			Email: "hogehoge@sample.com",
		},
		Profile: &model.Profile{
			Username: "sample_user",
			Image:    "http://xxxxxxxx.com",
		},
	}, true, nil
}
