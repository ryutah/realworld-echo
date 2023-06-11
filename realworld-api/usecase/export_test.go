package usecase

import "context"

// export for error_handler.go

var (
	WithErrorRendrer = withErrorRendrer
	BadRequest       = badRequest
	NotFound         = notFound
)

type ExportErrorHandler = errorHandler

func (e *ExportErrorHandler) Handle(ctx context.Context, err error, opts ...ErrorHandlerOption) error {
	return e.handle(ctx, err, opts...)
}

type ExportErrorHandlerConfig = errorHandlerConfig

func (e *ExportErrorHandlerConfig) AddRendrer(target error, f renderFunc) {
	e.rendrers = append(e.rendrers, errorRenderer{
		target:   target,
		renderer: f,
	})
}
