package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

type User interface {
	Get(context.Context, model.UserID) (*model.User, error)
	List(context.Context, ...model.UserID) ([]model.User, error)
}
