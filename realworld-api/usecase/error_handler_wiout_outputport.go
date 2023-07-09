package usecase

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
)

type ErrorHandlerWithoutOutputPortErrorHandleOption func(context.Context, error) (*FailResult, bool)

type ErrorHandlerWithoutOutputPort[R any] struct {
	errorReporter ErrorReporter
}

func NewErrorHandlerWithoutOutputPort[R any](reporter ErrorReporter) *ErrorHandlerWithoutOutputPort[R] {
	return &ErrorHandlerWithoutOutputPort[R]{
		errorReporter: reporter,
	}
}

func (e *ErrorHandlerWithoutOutputPort[R]) Handle(ctx context.Context, err error, opts ...ErrorHandlerWithoutOutputPortErrorHandleOption) *Result[R] {
	xlog.Warn(ctx, fmt.Sprintf("render error: %+v", err))

	for _, opt := range opts {
		if result, ok := opt(ctx, err); ok {
			return NewResultAsFailed[R](result)
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
	return NewResultAsFailed[R](newFaileResult(FailTypeInternalError, err))
}

func HandleAsBadRequest(targets ...error) ErrorHandlerWithoutOutputPortErrorHandleOption {
	return func(ctx context.Context, err error) (*FailResult, bool) {
		if includeInErrors(err, targets...) {
			xlog.Debug(ctx, "render bad request")
			return newFaileResult(FailTypeBadRequest, err), true
		}
		return nil, false
	}
}

func HandleAsNotFound(targets ...error) ErrorHandlerWithoutOutputPortErrorHandleOption {
	return func(ctx context.Context, err error) (*FailResult, bool) {
		if includeInErrors(err, targets...) {
			xlog.Debug(ctx, "render not found")
			return newFaileResult(FailTypeNotFound, err), true
		}
		return nil, false
	}
}

func newFaileResult(typ FailType, err error) *FailResult {
	return NewFailResult(
		typ,
		fmt.Sprintf("%v", err),
		errors.GetAllDetails(err)...,
	)
}

func includeInErrors(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
