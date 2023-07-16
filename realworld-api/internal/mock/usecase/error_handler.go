package usecase

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/stretchr/testify/mock"
)

var ErrorHandlerFuncNames = struct {
	Handle string
}{
	Handle: "Handle",
}

type MockErrorHandler[R any] struct {
	mock.Mock
}

var _ usecase.ErrorHandler[any] = (*MockErrorHandler[any])(nil)

func NewMockErrorHandler[R any]() *MockErrorHandler[R] {
	return &MockErrorHandler[R]{}
}

func (m *MockErrorHandler[R]) Handle(ctx context.Context, err error, opts ...usecase.ErrorHandlerOption) *usecase.Result[R] {
	args := m.Called(ctx, err, opts)
	return args.Get(0).(*usecase.Result[R])
}

func (m *MockErrorHandler[R]) Funcs() {
}
