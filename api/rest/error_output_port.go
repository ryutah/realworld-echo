package rest

import (
	"context"
	"net/http"

	"github.com/ryutah/realworld-echo/api/rest/gen"
	"github.com/ryutah/realworld-echo/usecase"
)

type errorOutputPort struct{}

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
