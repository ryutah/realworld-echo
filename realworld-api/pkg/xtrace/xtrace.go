package xtrace

import (
	"context"
	"net/http"
	"runtime"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/cockroachdb/errors"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type ProjectID string

type FinishTraceFunc func(context.Context) error

type Initializer interface {
	HandlerWithTracing(http.Handler) (http.Handler, FinishTraceFunc, error)
}

type stdoutTracingInitializer struct {
	sampler sdktrace.Sampler
}

func NewStdoutTracingInitializer(sampler sdktrace.Sampler) Initializer {
	return &stdoutTracingInitializer{
		sampler: sampler,
	}
}

func (s *stdoutTracingInitializer) HandlerWithTracing(h http.Handler) (http.Handler, FinishTraceFunc, error) {
	exporter, err := stdouttrace.New()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate exporter")
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(s.sampler),
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

type googleCloudTracingInitializer struct {
	projectID ProjectID
	sampler   sdktrace.Sampler
}

func NewGoogleCloudTracingInitializer(projectID ProjectID, sampler sdktrace.Sampler) Initializer {
	return &googleCloudTracingInitializer{
		projectID: projectID,
		sampler:   sampler,
	}
}

// examples
//   - https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/net/http/otelhttp/example/server/server.go
//   - https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/blob/main/example/trace/http/server/server.go
func (g *googleCloudTracingInitializer) HandlerWithTracing(h http.Handler) (http.Handler, FinishTraceFunc, error) {
	ctx := context.Background()
	exporter, err := texporter.New(texporter.WithProjectID(string(g.projectID)))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate exporter")
	}
	res, err := resource.New(
		ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("app_name"),
		),
	)
	if err != nil {
		return nil, nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(g.sampler),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			gcppropagator.CloudTraceOneWayPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	))

	return otelhttp.NewHandler(h, "server"), func(ctx context.Context) error {
		return tp.Shutdown(ctx)
	}, nil
}

func StartSpan(ctx context.Context) (context.Context, trace.Span) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("could not get caller function")
	}
	fn := runtime.FuncForPC(pc)
	// TODO: set trace name
	return otel.GetTracerProvider().Tracer("my_app").Start(ctx, fn.Name())
}
