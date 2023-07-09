package di

import (
	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/onmemory"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

func inputPortProvider() fx.Option {
	return fx.Provide(
		usecase.NewErrorHandler[article.GetArticleResult],
		article.NewGetArticle,
	)
}

func errorReportProvider() fx.Option {
	return fx.Provide(
		func() xerrorreport.Service {
			return xerrorreport.Service(config.GetConfig().Service)
		},
		func() xerrorreport.Version {
			return xerrorreport.Version(config.GetConfig().Version)
		},
		fx.Annotate(
			xerrorreport.NewErrorReporter,
			fx.As(new(usecase.ErrorReporter)),
		),
	)
}

func restProvider() fx.Option {
	return fx.Provide(
		rest.NewArticle,
		rest.NewServer,
		rest.NewExecuter,
	)
}

func localRepositoryProvider() fx.Option {
	return fx.Provide(
		onmemory.NewArticle,
	)
}

func localTraceProvider() fx.Option {
	return fx.Provide(
		sdktrace.NeverSample,
		xtrace.NewStdoutTracingInitializer,
	)
}

func InjectLocal(f func(e *rest.Extcuter)) *fx.App {
	return fx.New(
		localRepositoryProvider(),
		localTraceProvider(),
		errorReportProvider(),
		inputPortProvider(),
		restProvider(),
		fx.Invoke(f),
	)
}

func InjectAppEngine(f func(e *rest.Extcuter)) *fx.App {
	return fx.New(
		fx.Provide(),
	)
}

func InjectCloudRun(f func(e *rest.Extcuter)) *fx.App {
	return fx.New(
		fx.Provide(),
	)
}
