package usecase

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
)

type GenericsErrorReporter interface {
	Report(ctx context.Context, err error, errContext xerrorreport.ErrorContext)
}

type GenericsErrorResult struct {
	Message      string
	Descriptions []string
}

type GenericsErrorOutputPort[Ret any] interface {
	InternalError(context.Context, ErrorResult) (Ret, error)
	NotFound(context.Context, ErrorResult) (Ret, error)
	BadRequest(context.Context, ErrorResult) (Ret, error)
}

type GenericsErrorHandlerOption[Ret any] func(*genericsErrorHandlerConfig[Ret])

type GenericsErrorHandler[Ret any] interface {
	Handle(context.Context, error, ...GenericsErrorHandlerOption[Ret]) (Ret, error)
}

type genericsErrorHandler[Ret any] struct {
	errorReporter GenericsErrorReporter
	outputPort    GenericsErrorOutputPort[Ret]
}

func NewGenericsErrorHandler[Ret any](errorReporter GenericsErrorReporter, outputPort GenericsErrorOutputPort[Ret]) *genericsErrorHandler[Ret] {
	return &genericsErrorHandler[Ret]{
		errorReporter: errorReporter,
		outputPort:    outputPort,
	}
}

func (e *genericsErrorHandler[Ret]) Handle(ctx context.Context, err error, opts ...GenericsErrorHandlerOption[Ret]) (Ret, error) {
	xlog.Warn(ctx, fmt.Sprintf("render error: %+v", err))

	var opt genericsErrorHandlerConfig[Ret]
	for _, o := range opts {
		o(&opt)
	}

	for _, catch := range opt.rendrers {
		if errors.Is(err, catch.target) {
			return catch.renderer(ctx, e.outputPort, err)
		}
	}

	xlog.Debug(ctx, "render internal error")

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

type genericsRenderFunc[Ret any] func(context.Context, GenericsErrorOutputPort[Ret], error) (Ret, error)

type genericsErrorHandlerConfig[Ret any] struct {
	rendrers []genericsErrorRenderer[Ret]
}

type genericsErrorRenderer[Ret any] struct {
	target   error
	renderer genericsRenderFunc[Ret]
}

func WithGenericsErrorRendrer[Ret any](target error, f genericsRenderFunc[Ret]) GenericsErrorHandlerOption[Ret] {
	return func(opt *genericsErrorHandlerConfig[Ret]) {
		opt.rendrers = append(opt.rendrers, genericsErrorRenderer[Ret]{
			target:   target,
			renderer: f,
		})
	}
}

func GenericsBadRequest[Ret any](ctx context.Context, port GenericsErrorOutputPort[Ret], err error) (Ret, error) {
	xlog.Debug(ctx, "render bad request")
	return port.BadRequest(ctx, ErrorResult{
		Message:      fmt.Sprintf("%v", err),
		Descriptions: errors.GetAllDetails(err),
	})
}

func GenericsNotFound[Ret any](ctx context.Context, port GenericsErrorOutputPort[Ret], err error) (Ret, error) {
	xlog.Debug(ctx, "render not found")
	return port.NotFound(ctx, ErrorResult{
		Message:      fmt.Sprintf("%v", err),
		Descriptions: errors.GetAllDetails(err),
	})
}
