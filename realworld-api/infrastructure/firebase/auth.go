package firebase

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
)

type Auth struct{}

var _ service.Auth = (*Auth)(nil)

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) CurrentUser(_ context.Context) (*model.User, error) {
	return nil, errors.WithStack(derrors.Errors.NotAuthorized.Err)
}
