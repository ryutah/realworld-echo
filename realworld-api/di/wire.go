//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	restSet = wire.NewSet(
		rest.NewServer,
		rest.NewArticle,
		inputPortSet,
	)
	inputPortSet = wire.NewSet(
		article.NewArticle,
		usecase.NewErrorHandler,
		wire.Bind(new(article.GetArticleInputPort), new(*article.Article)),
		outputPortSet,
		repositorySet,
	)
	outputPortSet = wire.NewSet(
		rest.NewErrorOutputPort,
		rest.NewGetArticleOutputPort,
	)
	// TODO: set real repository
	repositorySet = wire.NewSet(
		wire.InterfaceValue(new(repository.Article), repository.Article(nil)),
	)
)

var localTraceInitializerSet = wire.NewSet(
	xtrace.NewStdoutTracingInitializer,
	// disaable sampling because it is too noisy for local development
	sdktrace.NeverSample,
	xerrorreport.NewErrorReporter,
	wire.Bind(new(usecase.ErrorReporter), new(*xerrorreport.ErrorReporter)),
)

var gcpTraceInitializerSet = wire.NewSet(
	xtrace.NewGoogleCloudTracingInitializer,
	sdktrace.AlwaysSample,
	xerrorreport.NewErrorReporter,
	wire.Bind(new(usecase.ErrorReporter), new(*xerrorreport.ErrorReporter)),
)

func InitializeLocalRestExecuter(service xerrorreport.Service, version xerrorreport.Version) *rest.Extcuter {
	panic(wire.Build(
		rest.NewExecuter,
		restSet,
		localTraceInitializerSet,
	))
}

func InitializeAppEngineRestExecuter(projectID xtrace.ProjectID, service xerrorreport.Service, version xerrorreport.Version) *rest.Extcuter {
	panic(wire.Build(
		rest.NewExecuter,
		restSet,
		gcpTraceInitializerSet,
	))
}

func InitializeCloudRunRestExecuter(projectID xtrace.ProjectID, service xerrorreport.Service, version xerrorreport.Version) *rest.Extcuter {
	panic(wire.Build(
		rest.NewExecuter,
		restSet,
		gcpTraceInitializerSet,
	))
}
