package testadata

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/internal/operations"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"go.uber.org/fx"
)

type (
	SampleCreateInputPort interface {
		Create(context.Context, SampleCreateResult) usecase.Result[SampleCreateResult]
	}
	SampleCreateParams struct{}
	SampleCreateResult struct{}
)

type SampleCreate struct {
	errorHandler usecase.ErrorHandler[SampleCreateResult]
	repository   struct{}
	service      struct{}
}

type SampleCreateDependencies struct {
	fx.In
	ErrorHandler usecase.ErrorHandler[SampleCreateResult]
}

func NewSampleCreate(d SampleCreateDependencies) *SampleCreate {
	return &SampleCreate{
		errorHandler: d.ErrorHandler,
		repository:   struct{}{},
		service:      struct{}{},
	}
}

func (s *SampleCreate) Create(ctx context.Context, param SampleCreateParams) usecase.Result[SampleCreateResult] {
	ctx, finish := operations.StartFunc(ctx, operations.FuncParam(param))
	defer finish()

	return usecase.Result[SampleCreateResult]{}
}
