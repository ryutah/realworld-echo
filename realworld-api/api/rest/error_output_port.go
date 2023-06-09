package rest

import (
	"context"
	"net/http"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type errorOutputPort struct {
	usecase.ErrorOutputPort
}

func NewErrorOutputPort() usecase.ErrorOutputPort {
	return &errorOutputPort{}
}

func (e *errorOutputPort) InternalError(ctx context.Context, result usecase.ErrorResult) error {
	c := echoContextFromContext(ctx)
	return c.JSON(http.StatusInternalServerError, gen.GenericError{
		Errors: struct {
			Body []string `json:"body"`
		}{
			Body: []string{
				result.Message,
			},
		},
	})
}

func (e *errorOutputPort) NotFound(_ context.Context, _ usecase.ErrorResult) error {
	panic("not implemented") // TODO: Implement
}

type genericsErrorOutputPort[Ret any] struct {
	usecase.GenericsErrorOutputPort[Ret]
}

func newGenericsErrorOutputPort[Ret any]() usecase.GenericsErrorOutputPort[Ret] {
	return &genericsErrorOutputPort[Ret]{}
}
