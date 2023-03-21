package rest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	. "github.com/ryutah/realworld-echo/api/rest"
	"github.com/ryutah/realworld-echo/internal/xtest"
)

type dummyEchoContext struct {
	echo.Context
}

func (d *dummyEchoContext) Request() *http.Request {
	dummyReq, _ := http.NewRequest(http.MethodGet, "/", nil)
	return dummyReq
}

func Test_Context(t *testing.T) {
	ec := new(dummyEchoContext)
	ctx := NewContext(ec)
	got := EchoContextFromContext(ctx)
	if diff := cmp.Diff(ec, got); diff != "" {
		t.Errorf(xtest.ErrorMsg.Diff, diff)
	}
}

func Test_Context_ShouldBePanic(t *testing.T) {
	defer func() {
		want := "context has no echo.Context"
		err := recover()
		if diff := cmp.Diff(want, err); diff != "" {
			t.Errorf(xtest.ErrorMsg.Diff, diff)
		}
	}()

	type dummyEchoContext struct {
		echo.Context
	}

	ctx := context.Background()
	_ = EchoContextFromContext(ctx)
}
