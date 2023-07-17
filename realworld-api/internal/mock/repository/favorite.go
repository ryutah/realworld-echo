package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/stretchr/testify/mock"
)

var FavroteFuncNames = struct {
	ListBySlug string
}{
	ListBySlug: "ListBySlug",
}

type MockFavorite struct {
	mock.Mock
}

var _ repository.Favorite = (*MockFavorite)(nil)

func NewMockFavorite() *MockFavorite {
	return &MockFavorite{}
}

func (m *MockFavorite) ListBySlug(ctx context.Context, slug model.Slug) (model.FavoriteSlice, error) {
	args := m.Called(ctx, slug)
	return args.Get(0).(model.FavoriteSlice), args.Error(1)
}
