package usecase

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
)

type ErrorReporter interface {
	Report(ctx context.Context, err error, errContext xerrorreport.ErrorContext)
}

type ErrorHandlerOption interface {
	Apply(*ErrorHandlerConfig)
}

// type ErrorHandlerOption func(*errorHandlerConfig)

type errorHandlerFunc func(context.Context, error) (*FailResult, bool)

type ErrorHandlerConfig struct {
	handlers []errorHandlerFunc
}

func (e *ErrorHandlerConfig) addHandlers(opt errorHandlerFunc) {
	e.handlers = append(e.handlers, opt)
}

type ErrorHandler[R any] interface {
	Handle(context.Context, error, ...ErrorHandlerOption) *Result[R]
}

type errorHandler[R any] struct {
	errorReporter ErrorReporter
	service       struct {
		auth service.Auth
	}
}

func NewErrorHandler[R any](reporter ErrorReporter, authService service.Auth) ErrorHandler[R] {
	return &errorHandler[R]{
		errorReporter: reporter,
		service: struct {
			auth service.Auth
		}{
			auth: authService,
		},
	}
}

func (e *errorHandler[R]) Handle(ctx context.Context, err error, opts ...ErrorHandlerOption) *Result[R] {
	xlog.Warn(ctx, fmt.Sprintf("render error: %+v", err))

	var config ErrorHandlerConfig
	for _, opt := range opts {
		opt.Apply(&config)
	}

	for _, h := range config.handlers {
		if result, ok := h(ctx, err); ok {
			return Fail[R](result)
		}
	}

	var uid string
	if user, err := e.service.auth.CurrentUser(ctx); err == nil {
		uid = user.ID.String()
	}
	xlog.Debug(ctx, "render internal error")
	file, line, fn, _ := errors.GetOneLineSource(err)
	e.errorReporter.Report(ctx, err, xerrorreport.ErrorContext{
		User: uid,
		Location: xerrorreport.Location{
			File:     file,
			Line:     line,
			Function: fn,
		},
	})
	return Fail[R](newFaileResult(FailTypeInternalError, err))
}

type badRequestHandler struct {
	targets []error
}

func (h *badRequestHandler) Apply(c *ErrorHandlerConfig) {
	c.addHandlers(func(ctx context.Context, err error) (*FailResult, bool) {
		if includeInErrors(err, h.targets...) {
			xlog.Debug(ctx, "render bad request")
			return newFaileResult(FailTypeBadRequest, err), true
		}
		return nil, false
	})
}

func WithBadRequestHandler(targets ...error) ErrorHandlerOption {
	return &badRequestHandler{
		targets: targets,
	}
}

type notFoundHandler struct {
	targets []error
}

func (h *notFoundHandler) Apply(c *ErrorHandlerConfig) {
	c.addHandlers(func(ctx context.Context, err error) (*FailResult, bool) {
		if includeInErrors(err, h.targets...) {
			xlog.Debug(ctx, "render not found")
			return newFaileResult(FailTypeNotFound, err), true
		}
		return nil, false
	})
}

func WithNotFoundHandler(targets ...error) ErrorHandlerOption {
	return &notFoundHandler{
		targets: targets,
	}
}

type unauthorizedHandler struct {
	targets []error
}

func (h *unauthorizedHandler) Apply(c *ErrorHandlerConfig) {
	c.addHandlers(func(ctx context.Context, err error) (*FailResult, bool) {
		if includeInErrors(err, h.targets...) {
			xlog.Debug(ctx, "render unauthorized")
			return newFaileResult(FailTypeUnauthorized, err), true
		}
		return nil, false
	})
}

func WithUnauthorizedHandler(targets ...error) ErrorHandlerOption {
	return &unauthorizedHandler{
		targets: targets,
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
