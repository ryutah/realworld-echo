package middleware

import (
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/pkg/xlog"
	"go.uber.org/zap"
)

func WithLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := xlog.NewLogger()

		req := c.Request()
		spanCtx, err := gcppropagator.SpanContextFromRequest(req)
		if err == nil {
			logger = logger.
				With(zap.String("logging.googleapis.com/trace", spanCtx.TraceID().String())).
				With(zap.String("logging.googleapis.com/spanId", spanCtx.SpanID().String())).
				With(zap.Bool("logging.googleapis.com/trace_sampled", spanCtx.IsSampled()))
		}

		newReq := req.WithContext(xlog.ContextWithLogger(req.Context(), logger))
		c.SetRequest(newReq)

		return next(c)
	}
}
