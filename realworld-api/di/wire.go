//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	restSet = wire.NewSet(
		rest.NewServer,
		rest.NewArticle,
		inputPortSet,
	)
	inputPortSet = wire.NewSet(
		usecase.NewArticle,
		wire.Bind(new(usecase.GetArticleInputPort), new(*usecase.Article)),
		outputPortSet,
	)
	outputPortSet = wire.NewSet(
		rest.NewErrorOutputPort,
		rest.NewGetArticleOutputPort,
	)
)

var localTraceInitializerSet = wire.NewSet(
	xtrace.NewStdoutTracingInitializer,
	// disaable sampling because it is too noisy for local development
	sdktrace.NeverSample,
)

var gcpTraceInitializerSet = wire.NewSet(
	xtrace.NewGoogleCloudTracingInitializer,
	sdktrace.AlwaysSample,
)

func InitializeLocalRestExecuter() *rest.Extcuter {
	panic(wire.Build(
		rest.NewExecuter,
		restSet,
		localTraceInitializerSet,
	))
}

func InitializeAppEngineRestExecuter(projectID string) *rest.Extcuter {
	panic(wire.Build(
		rest.NewExecuter,
		restSet,
		gcpTraceInitializerSet,
	))
}
