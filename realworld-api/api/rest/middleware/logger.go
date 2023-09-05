package middleware

import (
	"fmt"
	"log/slog"

	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xslog"
)

func WithLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := xslog.NewLogger()

		req := c.Request()
		spanCtx, err := gcppropagator.SpanContextFromRequest(req)
		if err == nil {
			traceID := fmt.Sprintf("projects/%s/traces/%s", config.GetConfig().ProjectID, spanCtx.TraceID().String())
			logger = logger.
				With(slog.String("logging.googleapis.com/trace", traceID)).
				With(slog.String("logging.googleapis.com/spanId", spanCtx.SpanID().String())).
				With(slog.Bool("logging.googleapis.com/trace_sampled", spanCtx.IsSampled()))
		}

		newReq := req.WithContext(xslog.ContextWithLogger(req.Context(), logger))
		c.SetRequest(newReq)

		return next(c)
	}
}
