package usecase

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
)

type ErrorReporter interface {
	Report(ctx context.Context, err error, errContext xerrorreport.ErrorContext)
}

type ErrorResult struct {
	Message      string
	Descriptions []string
}

type ErrorOutputPort interface {
	InternalError(context.Context, ErrorResult) error
	NotFound(context.Context, ErrorResult) error
	BadRequest(context.Context, ErrorResult) error
}

type ErrorHandlerOption func(*errorHandlerConfig)

type ErrorHandler interface {
	handle(context.Context, error, ...ErrorHandlerOption) error
}

type errorHandler struct {
	errorReporter ErrorReporter
	outputPort    ErrorOutputPort
}

func NewErrorHandler(errorReporter ErrorReporter, outputPort ErrorOutputPort) ErrorHandler {
	return &errorHandler{
		errorReporter: errorReporter,
		outputPort:    outputPort,
	}
}

func (e *errorHandler) handle(ctx context.Context, err error, opts ...ErrorHandlerOption) error {
	var opt errorHandlerConfig
	for _, o := range opts {
		o(&opt)
	}

	for _, catch := range opt.rendrers {
		if errors.Is(err, catch.target) {
			return catch.renderer(ctx, e.outputPort, err)
		}
	}

	file, line, fn, _ := errors.GetOneLineSource(err)
	e.errorReporter.Report(ctx, err, xerrorreport.ErrorContext{
		User: "", // TODO: should be get user id from context
		Location: xerrorreport.Location{
			File:     file,
			Line:     line,
			Function: fn,
		},
	})
	return e.outputPort.InternalError(ctx, ErrorResult{
		Message:      err.Error(),
		Descriptions: errors.GetAllDetails(err),
	})
}

type renderFunc func(context.Context, ErrorOutputPort, error) error

type errorHandlerConfig struct {
	rendrers []errorRenderer
}

type errorRenderer struct {
	target   error
	renderer renderFunc
}

func withErrorRendrer(target error, f renderFunc) ErrorHandlerOption {
	return func(opt *errorHandlerConfig) {
		opt.rendrers = append(opt.rendrers, errorRenderer{
			target:   target,
			renderer: f,
		})
	}
}

func badRequest(ctx context.Context, port ErrorOutputPort, err error) error {
	return port.BadRequest(ctx, ErrorResult{
		Message:      fmt.Sprintf("%v", err),
		Descriptions: errors.GetAllDetails(err),
	})
}

func notFound(ctx context.Context, port ErrorOutputPort, err error) error {
	return port.NotFound(ctx, ErrorResult{
		Message:      fmt.Sprintf("%v", err),
		Descriptions: errors.GetAllDetails(err),
	})
}
