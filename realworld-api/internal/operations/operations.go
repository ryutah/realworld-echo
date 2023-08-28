package operations

import (
	"context"
	"fmt"
	"runtime"

	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"go.uber.org/zap"
)

func StartFunc(ctx context.Context, fields ...zap.Field) (context.Context, func()) {
	ctx, span := xtrace.StartSpan(ctx)
	ctx = xlog.ContextWithLogFields(ctx, fields...)

	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		xlog.Info(ctx, fmt.Sprintf("start %v", fn.Name()), fields...)
	}

	return ctx, func() {
		span.End()
		if fn != nil {
			xlog.Info(ctx, fmt.Sprintf("end %v", fn.Name()))
		}
	}
}
