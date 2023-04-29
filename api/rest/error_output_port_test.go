package rest_test

import (
	"context"
	"testing"

	. "github.com/ryutah/realworld-echo/api/rest"
	"github.com/ryutah/realworld-echo/usecase"
)

func Test_ErrorOutputPort_InternalError(t *testing.T) {
	type args struct {
		ctx    context.Context
		result usecase.ErrorResult
	}
	tests := []struct {
		name    string
		e       *ErrorOutputPort
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// e := &ErrorOutputPort{}
			// if err := e.InternalError(tt.args.ctx, tt.args.result); (err != nil) != tt.wantErr {
			// 	t.Errorf("errorOutputPort.InternalError() error = %v, wantErr %v", err, tt.wantErr)
			// }
		})
	}
}

func Test_ErrorOutputPort_NotFound(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 usecase.ErrorResult
	}
	tests := []struct {
		name    string
		e       *ErrorOutputPort
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// e := &ErrorOutputPort{}
			// if err := e.NotFound(tt.args.in0, tt.args.in1); (err != nil) != tt.wantErr {
			// 	t.Errorf("errorOutputPort.NotFound() error = %v, wantErr %v", err, tt.wantErr)
			// }
		})
	}
}
