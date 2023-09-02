package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	mock_auth_service "github.com/ryutah/realworld-echo/realworld-api/internal/mock/auth/service"
	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/usecase"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ErrorHandler_Handle(t *testing.T) {
	dummyError := errors.WithDetail(errors.New("dummy error"), "dummy_error_detail")

	type mocks_errorReporter struct {
		report_args_error error
	}
	type mocks_authService struct {
		currentUser_returns_user  *model.User
		currentUser_returns_error error
	}
	type mocks_errorHandlerOption struct {
		apply_params_errorHandlerFunc ErrorHandlerFunc
	}
	type mocks struct {
		errorReporter       mocks_errorReporter
		errorHandlerOptions []mocks_errorHandlerOption
		authService         mocks_authService
	}
	type configs struct {
		isInternalError bool
	}
	type args struct {
		err error
	}
	type wants struct {
		result *Result[any]
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		configs configs
		wants   wants
	}{
		{
			name: "when_given_error_match_custome_error_handler_error_should_call_custom_error_handler_and_return_expected_result",
			args: args{
				err: dummyError,
			},
			mocks: mocks{
				errorHandlerOptions: []mocks_errorHandlerOption{
					{
						apply_params_errorHandlerFunc: func(ctx context.Context, err error) (*FailResult, bool) {
							return NewFailResult(FailTypeBadRequest, "test_faile_result"), true
						},
					},
				},
			},
			configs: configs{
				isInternalError: false,
			},
			wants: wants{
				result: Fail[any](NewFailResult(FailTypeBadRequest, "test_faile_result")),
			},
		},
		{
			name: "when_given_second_option_error_match_custome_error_handler_error_should_call_custom_error_handler_and_return_expected_result",
			args: args{
				err: dummyError,
			},
			mocks: mocks{
				errorHandlerOptions: []mocks_errorHandlerOption{
					{
						apply_params_errorHandlerFunc: func(ctx context.Context, err error) (*FailResult, bool) {
							return nil, false
						},
					},
					{
						apply_params_errorHandlerFunc: func(ctx context.Context, err error) (*FailResult, bool) {
							return NewFailResult(FailTypeBadRequest, "test_faile_result"), true
						},
					},
					{
						apply_params_errorHandlerFunc: func(ctx context.Context, err error) (*FailResult, bool) {
							return nil, false
						},
					},
				},
			},
			configs: configs{
				isInternalError: false,
			},
			wants: wants{
				result: Fail[any](NewFailResult(FailTypeBadRequest, "test_faile_result")),
			},
		},
		{
			name: "when_given_error_not_match_custome_error_handler_error_should_call_ErroReporter#Report_and_ErrorOutputPort#InternalError_and_return_expected_result",
			args: args{
				err: dummyError,
			},
			mocks: mocks{
				errorReporter: mocks_errorReporter{
					report_args_error: dummyError,
				},
				authService: mocks_authService{
					currentUser_returns_user: &model.User{
						ID: "user_id",
					},
				},
				errorHandlerOptions: []mocks_errorHandlerOption{
					{
						apply_params_errorHandlerFunc: func(ctx context.Context, err error) (*FailResult, bool) {
							return nil, false
						},
					},
				},
			},
			configs: configs{
				isInternalError: true,
			},
			wants: wants{
				result: Fail[any](
					NewFailResult(FailTypeInternalError, fmt.Sprintf("%v", dummyError), "dummy_error_detail"),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var errorHandlerOptions []*mock_usecase.MockErrorHandlerOption
			errorReporter := mock_usecase.NewMockErrorReporter(t)
			authService := mock_auth_service.NewMockAuth(t)

			for i := range tt.mocks.errorHandlerOptions {
				idx := i
				opt := mock_usecase.NewMockErrorHandlerOption(t)
				opt.EXPECT().
					Apply(mock.Anything).
					Run(func(c *ErrorHandlerConfig) {
						c.AddHandlers(tt.mocks.errorHandlerOptions[idx].apply_params_errorHandlerFunc)
					})
				errorHandlerOptions = append(errorHandlerOptions, opt)
			}

			if tt.configs.isInternalError {
				errorReporter.EXPECT().Report(
					mock.Anything,
					tt.mocks.errorReporter.report_args_error,
					mock.Anything,
				)
				authService.EXPECT().
					CurrentUser(mock.Anything).
					Return(
						tt.mocks.authService.currentUser_returns_user,
						tt.mocks.authService.currentUser_returns_error,
					)
			}

			got := NewErrorHandler[any](errorReporter, authService).Handle(
				context.Background(),
				tt.args.err,
				lo.Map(
					errorHandlerOptions,
					func(i *mock_usecase.MockErrorHandlerOption, _ int) ErrorHandlerOption {
						return ErrorHandlerOption(i)
					},
				)...,
			)
			assert.Equal(t, tt.wants.result, got)
			errorReporter.AssertExpectations(t)
		})
	}
}

func Test_WithNotFoundHandler(t *testing.T) {
	dummyError := errors.WithDetail(errors.New("dummy"), "dummy_error_detail_1")
	dummyError2 := errors.WithDetail(errors.New("dummy2"), "dummy_error2_detail_1")

	type args struct {
		errs []error
	}
	type calls struct {
		err error
	}
	type wants struct {
		result *FailResult
		ok     bool
	}

	tests := []struct {
		name  string
		args  args
		calls calls
		wants wants
	}{
		{
			name: "when_given_any_errors_should_return_expected_result_and_true",
			args: args{
				errs: []error{dummyError, dummyError2},
			},
			calls: calls{
				err: dummyError,
			},
			wants: wants{
				result: NewFailResult(FailTypeNotFound, fmt.Sprintf("%v", dummyError), "dummy_error_detail_1"),
				ok:     true,
			},
		},
		{
			name: "when_given_nil_as_error_should_return_nil_and_false",
			args: args{
				errs: []error{nil},
			},
			calls: calls{
				err: dummyError,
			},
			wants: wants{
				result: nil,
				ok:     false,
			},
		},
		{
			name: "when_given_not_match_error_should_return_nil_and_false",
			args: args{
				errs: []error{dummyError},
			},
			calls: calls{
				err: dummyError2,
			},
			wants: wants{
				result: nil,
				ok:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := WithNotFoundHandler(tt.args.errs...)
			var config ErrorHandlerConfig
			opt.Apply(&config)

			got, ok := config.Handlers()[0](context.Background(), tt.calls.err)
			assert.Equal(t, tt.wants.result, got, "result")
			assert.Equal(t, tt.wants.ok, ok, "ok")
		})
	}
}

func Test_WithBadRequestHandler(t *testing.T) {
	dummyError := errors.WithDetail(errors.New("dummy"), "dummy_error_detail_1")
	dummyError2 := errors.WithDetail(errors.New("dummy2"), "dummy_error2_detail_1")

	type args struct {
		errs []error
	}
	type calls struct {
		err error
	}
	type wants struct {
		result *FailResult
		ok     bool
	}

	tests := []struct {
		name  string
		args  args
		calls calls
		wants wants
	}{
		{
			name: "when_given_any_errors_should_return_expected_result_and_true",
			args: args{
				errs: []error{dummyError, dummyError2},
			},
			calls: calls{
				err: dummyError,
			},
			wants: wants{
				result: NewFailResult(FailTypeBadRequest, fmt.Sprintf("%v", dummyError), "dummy_error_detail_1"),
				ok:     true,
			},
		},
		{
			name: "when_given_nil_as_error_should_return_nil_and_false",
			args: args{
				errs: []error{nil},
			},
			calls: calls{
				err: dummyError,
			},
			wants: wants{
				result: nil,
				ok:     false,
			},
		},
		{
			name: "when_given_not_match_error_should_return_nil_and_false",
			args: args{
				errs: []error{dummyError},
			},
			calls: calls{
				err: dummyError2,
			},
			wants: wants{
				result: nil,
				ok:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := WithBadRequestHandler(tt.args.errs...)
			var config ErrorHandlerConfig
			opt.Apply(&config)

			got, ok := config.Handlers()[0](context.Background(), tt.calls.err)
			assert.Equal(t, tt.wants.result, got, "result")
			assert.Equal(t, tt.wants.ok, ok, "ok")
		})
	}
}

func Test_WithUnauthorizedHandler(t *testing.T) {
	dummyError := errors.WithDetail(errors.New("dummy"), "dummy_error_detail_1")
	dummyError2 := errors.WithDetail(errors.New("dummy2"), "dummy_error2_detail_1")

	type args struct {
		errs []error
	}
	type calls struct {
		err error
	}
	type wants struct {
		result *FailResult
		ok     bool
	}

	tests := []struct {
		name  string
		args  args
		calls calls
		wants wants
	}{
		{
			name: "when_given_any_errors_should_return_expected_result_and_true",
			args: args{
				errs: []error{dummyError, dummyError2},
			},
			calls: calls{
				err: dummyError,
			},
			wants: wants{
				result: NewFailResult(FailTypeUnauthorized, fmt.Sprintf("%v", dummyError), "dummy_error_detail_1"),
				ok:     true,
			},
		},
		{
			name: "when_given_nil_as_error_should_return_nil_and_false",
			args: args{
				errs: []error{nil},
			},
			calls: calls{
				err: dummyError,
			},
			wants: wants{
				result: nil,
				ok:     false,
			},
		},
		{
			name: "when_given_not_match_error_should_return_nil_and_false",
			args: args{
				errs: []error{dummyError},
			},
			calls: calls{
				err: dummyError2,
			},
			wants: wants{
				result: nil,
				ok:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := WithUnauthorizedHandler(tt.args.errs...)
			var config ErrorHandlerConfig
			opt.Apply(&config)

			got, ok := config.Handlers()[0](context.Background(), tt.calls.err)
			assert.Equal(t, tt.wants.result, got, "result")
			assert.Equal(t, tt.wants.ok, ok, "ok")
		})
	}
}
