package usecase

import "context"

type ErrorResult struct {
	Message      string
	Descriptions []string
}

type ErrorOutputPort interface {
	InternalError(context.Context, ErrorResult) error
	NotFound(context.Context, ErrorResult) error
}

type errorHandler struct {
	outputPort ErrorOutputPort
}

func (e *errorHandler) handleError(ctx context.Context, err error) error {
	return e.outputPort.InternalError(ctx, ErrorResult{
		Message:      "",
		Descriptions: []string{},
	})
}
