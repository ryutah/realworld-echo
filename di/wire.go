//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ryutah/realworld-echo/api/rest"
	"github.com/ryutah/realworld-echo/pkg/xtrace"
	"github.com/ryutah/realworld-echo/usecase"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	localRestSet = wire.NewSet(
		rest.NewServer,
		rest.NewArticle,
		localInputPortSet,
	)
	localInputPortSet = wire.NewSet(
		usecase.NewArticle,
		wire.Bind(new(usecase.GetArticleInputPort), new(*usecase.Article)),
		localOutputPortSet,
	)
	localOutputPortSet = wire.NewSet(
		rest.NewErrorOutputPort,
		rest.NewGetArticleOutputPort,
	)
	traceInitializerSet = wire.NewSet(
		xtrace.NewGoogleCloudTracingInitializer,
		sdktrace.NeverSample,
	)
)

func InitializeRestExecuter(projectID string) *rest.Extcuter {
	panic(wire.Build(
		rest.NewExecuter,
		localRestSet,
		traceInitializerSet,
	))
}
