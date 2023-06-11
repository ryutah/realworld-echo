package rest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtesting"
)

func Test_Context(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctxMock := NewMockContext(ctrl)
	ctxMock.EXPECT().Request().Times(1).Return(new(http.Request))

	ctx := NewContext(ctxMock)
	got := EchoContextFromContext(ctx)

	if diff := cmp.Diff(ctxMock, got, cmpopts.IgnoreUnexported(MockContext{})); diff != "" {
		xtesting.PrintDiff(t, "NewContext", diff)
	}
}

func Test_Context_ShouldBePanic(t *testing.T) {
	defer func() {
		want := "context has no echo.Context"
		err := recover()
		if diff := cmp.Diff(want, err); diff != "" {
			xtesting.PrintDiff(t, "panic", diff)
		}
	}()
	ctx := context.Background()
	_ = EchoContextFromContext(ctx)
}
