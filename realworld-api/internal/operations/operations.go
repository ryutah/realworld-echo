package operations

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"

	"github.com/ryutah/realworld-echo/realworld-api/pkg/xslog"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
)

func FuncParam(param any) slog.Attr {
	return slog.Any("param", param)
}

func StartFunc(ctx context.Context, fields ...slog.Attr) (context.Context, func()) {
	ctx, span := xtrace.StartSpan(ctx)
	ctx = xslog.ContextWithAttrs(ctx, fields...)

	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		xslog.Info(ctx, fmt.Sprintf("start %v", fn.Name()), fields...)
	}

	return ctx, func() {
		span.End()
		if fn != nil {
			xslog.Info(ctx, fmt.Sprintf("end %v", fn.Name()))
		}
	}
}
