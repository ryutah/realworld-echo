package xtrace

import (
	"context"
	"net/http"

	"github.com/cockroachdb/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type FinishTraceFunc func(context.Context) error

type Initializer interface {
	HandlerWithTracing(http.Handler) (http.Handler, FinishTraceFunc, error)
}

type googleCloudTracingInitializer struct {
	projectID string
	sampler   sdktrace.Sampler
}

func NewGoogleCloudTracingInitializer(projectID string, sampler sdktrace.Sampler) Initializer {
	return &googleCloudTracingInitializer{
		projectID: projectID,
		sampler:   sampler,
	}
}

// examples
//   - https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/net/http/otelhttp/example/server/server.go
//   - https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/blob/main/example/trace/http/server/server.go
func (g *googleCloudTracingInitializer) HandlerWithTracing(h http.Handler) (http.Handler, FinishTraceFunc, error) {
	exporter, err := stdouttrace.New()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate exporter")
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(g.sampler),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{},
	))
	return otelhttp.NewHandler(h, "server"), func(ctx context.Context) error {
		return tp.Shutdown(ctx)
	}, nil
}

func AddSpan(ctx context.Context, name string) trace.Span {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(name)
	return span
}
