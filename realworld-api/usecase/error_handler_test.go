package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/golang/mock/gomock"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtesting"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/mock"
)

func Test_RrrorHandler_Handle(t *testing.T) {
	dummyError := errors.WithDetail(errors.New("dummy error"), "dummy_error_detail")
	dummyError2 := errors.WithDetail(errors.New("dummy error2"), "dummy_error_detail2")

	type mocks_errorReporter struct {
		report_args_error error
	}
	type mocks_errorOutputPort struct {
		internalError_args_errorResult ErrorResult
		internalError_returns_error    error
	}
	type mocks struct {
		errorReporter mocks_errorReporter
		outputPort    mocks_errorOutputPort
	}
	type configs struct {
		isInternalError bool
	}
	type args struct {
		ctx  context.Context
		err  error
		opts []ErrorHandlerOption
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		configs configs
		wants   error
	}{
		{
			name: "when_given_error_match_custome_error_handler_error_should_call_custom_error_handler_and_return_nil",
			args: args{
				ctx: context.TODO(),
				err: dummyError,
				opts: []ErrorHandlerOption{
					func(ehc *ExportErrorHandlerConfig) {
						ehc.AddRendrer(dummyError, func(ctx context.Context, outputPort ErrorOutputPort, err error) error {
							return nil
						})
					},
				},
			},
			configs: configs{
				isInternalError: false,
			},
			wants: nil,
		},
		{
			name: "when_given_error_not_match_custome_error_handler_error_should_call_ErroReporter#Report_and_ErrorOutputPort#InternalError_and_return_nil",
			args: args{
				ctx: context.TODO(),
				err: dummyError2,
				opts: []ErrorHandlerOption{
					func(ehc *ExportErrorHandlerConfig) {
						ehc.AddRendrer(dummyError, func(ctx context.Context, outputPort ErrorOutputPort, err error) error {
							return nil
						})
					},
				},
			},
			mocks: mocks{
				errorReporter: mocks_errorReporter{
					report_args_error: dummyError2,
				},
				outputPort: mocks_errorOutputPort{
					internalError_args_errorResult: ErrorResult{
						Message:      fmt.Sprintf("%v", dummyError2),
						Descriptions: errors.GetAllDetails(dummyError2),
					},
					internalError_returns_error: nil,
				},
			},
			configs: configs{
				isInternalError: true,
			},
			wants: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			errorReporter := mock.NewMockErrorReporter(ctrl)
			outputPort := mock.NewMockErrorOutputPort(ctrl)

			if tt.configs.isInternalError {
				errorReporter.EXPECT().Report(gomock.Not(nil), tt.mocks.errorReporter.report_args_error, gomock.Any())
				outputPort.EXPECT().InternalError(gomock.Not(nil), tt.mocks.outputPort.internalError_args_errorResult).Return(tt.mocks.outputPort.internalError_returns_error)
			}

			e := NewErrorHandler(errorReporter, outputPort).(*ExportErrorHandler)
			err := e.Handle(tt.args.ctx, tt.args.err, tt.args.opts...)
			xtesting.CompareError(t, "handle", tt.wants, err)
		})
	}
}

func Test_NotFound(t *testing.T) {
	dummyError := errors.WithDetail(errors.New("dummy"), "dummy_error_detail")

	type mocks_errorOutputPort struct {
		notFound_args_errorResult ErrorResult
		notFound_returns_error    error
	}
	type mocks struct {
		errorOutputPort mocks_errorOutputPort
	}
	type args struct {
		ctx context.Context
		err error
	}
	tests := []struct {
		name  string
		args  args
		mocks mocks
		wants error
	}{
		{
			name: "when_given_any_error_should_call_ErrorOutputPort#NotFound_with_expected_args_and_return_nil",
			args: args{
				ctx: context.TODO(),
				err: errors.WithDetail(derrors.Errors.NotFound.Err, derrors.Errors.NotFound.Message),
			},
			mocks: mocks{
				errorOutputPort: mocks_errorOutputPort{
					notFound_args_errorResult: ErrorResult{
						Message: derrors.Errors.NotFound.Err.Error(),
						Descriptions: []string{
							derrors.Errors.NotFound.Message,
						},
					},
					notFound_returns_error: nil,
				},
			},
			wants: nil,
		},
		{
			name: "when_given_no_details_error_should_call_ErrorOutputPort#NotFound_with_expected_args_and_return_nil",
			args: args{
				ctx: context.TODO(),
				err: derrors.Errors.NotFound.Err,
			},
			mocks: mocks{
				errorOutputPort: mocks_errorOutputPort{
					notFound_args_errorResult: ErrorResult{
						Message: derrors.Errors.NotFound.Err.Error(),
					},
					notFound_returns_error: nil,
				},
			},
			wants: nil,
		},
		{
			name: "when_given_nil_as_error_should_call_ErrorOutputPort#NotFound_with_expected_args_and_return_nil",
			args: args{
				ctx: context.TODO(),
				err: nil,
			},
			mocks: mocks{
				errorOutputPort: mocks_errorOutputPort{
					notFound_args_errorResult: ErrorResult{
						Message: fmt.Sprintf("%v", nil),
					},
					notFound_returns_error: nil,
				},
			},
			wants: nil,
		},
		{
			name: "when_given_any_error_should_call_ErrorOutputPort#NotFound_returned_error_with_expected_args_and_return_error",
			args: args{
				ctx: context.TODO(),
				err: errors.WithDetail(derrors.Errors.NotFound.Err, derrors.Errors.NotFound.Message),
			},
			mocks: mocks{
				errorOutputPort: mocks_errorOutputPort{
					notFound_args_errorResult: ErrorResult{
						Message: derrors.Errors.NotFound.Err.Error(),
						Descriptions: []string{
							derrors.Errors.NotFound.Message,
						},
					},
					notFound_returns_error: dummyError,
				},
			},
			wants: dummyError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			outputPort := mock.NewMockErrorOutputPort(ctrl)
			outputPort.EXPECT().
				NotFound(tt.args.ctx, tt.mocks.errorOutputPort.notFound_args_errorResult).
				Return(tt.mocks.errorOutputPort.notFound_returns_error)

			err := NotFound(tt.args.ctx, outputPort, tt.args.err)

			xtesting.CompareError(t, "NotFound", tt.wants, err)
		})
	}
}
