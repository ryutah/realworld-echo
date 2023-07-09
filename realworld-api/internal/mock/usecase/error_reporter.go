package usecase

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/stretchr/testify/mock"
)

var ErrorReporterFuncNames = struct {
	Report string
}{
	Report: "Report",
}

type MockErrorReporter struct {
	mock.Mock
}

var _ (usecase.ErrorReporter) = (*MockErrorReporter)(nil)

func NewMockErrorReporter() *MockErrorReporter {
	return &MockErrorReporter{}
}

func (m *MockErrorReporter) Report(ctx context.Context, err error, errContext xerrorreport.ErrorContext) {
	m.Called(ctx, err, errContext)
}
