package service

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type Auth interface {
	CurrentUser(context.Context) (*model.User, error)
}
