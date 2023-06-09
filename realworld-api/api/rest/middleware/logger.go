package middleware

import (
	"fmt"

	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"go.uber.org/zap"
)

func WithLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := xlog.NewLogger()

		req := c.Request()
		spanCtx, err := gcppropagator.SpanContextFromRequest(req)
		if err == nil {
			traceID := fmt.Sprintf("projects/%s/traces/%s", config.GetConfig().ProjectID, spanCtx.TraceID().String())
			logger = logger.
				With(zap.String("logging.googleapis.com/trace", traceID)).
				With(zap.String("logging.googleapis.com/spanId", spanCtx.SpanID().String())).
				With(zap.Bool("logging.googleapis.com/trace_sampled", spanCtx.IsSampled()))
		}

		newReq := req.WithContext(xlog.ContextWithLogger(req.Context(), logger))
		c.SetRequest(newReq)

		return next(c)
	}
}
