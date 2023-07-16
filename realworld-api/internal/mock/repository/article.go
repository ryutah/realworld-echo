package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/stretchr/testify/mock"
)

var ArticleFuncNames = struct {
	GenerateID string
	Get        string
	Save       string
	Search     string
}{
	GenerateID: "GenerateID",
	Get:        "Get",
	Save:       "Save",
	Search:     "Search",
}

type MockArticle struct {
	mock.Mock
}

var _ (repository.Article) = &MockArticle{}

func NewMockArticle() *MockArticle {
	return &MockArticle{}
}

func (m *MockArticle) GenerateID(ctx context.Context) (model.Slug, error) {
	args := m.Called(ctx)
	return args.Get(0).(model.Slug), args.Error(1)
}

func (m *MockArticle) Get(ctx context.Context, slug model.Slug) (*model.Article, error) {
	args := m.Called(ctx, slug)
	return args.Get(0).(*model.Article), args.Error(1)
}

func (m *MockArticle) Save(ctx context.Context, article model.Article) error {
	args := m.Called(ctx, article)
	return args.Error(0)
}

func (m *MockArticle) Search(ctx context.Context, param repository.ArticleSearchParam) ([]model.Article, error) {
	args := m.Called(ctx, param)
	return args.Get(0).([]model.Article), args.Error(1)
}
