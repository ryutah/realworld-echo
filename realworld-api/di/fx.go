package di

import (
	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/firebase"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
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
		usecase.NewErrorHandler[article.ListArticleResult],
		article.NewGetArticle,
		article.NewListArticle,
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
		fx.Annotate(sqlc.NewDBManager, fx.As(new(sqlc.DBManager))),
		fx.Annotate(sqlc.NewArtile, fx.As(new(repository.Article))),
		fx.Annotate(sqlc.NewFavorite, fx.As(new(repository.Favorite))),
	)
}

func localAuthServiceProvider() fx.Option {
	return fx.Provide(
		fx.Annotate(firebase.NewAuth, fx.As(new(service.Auth))),
	)
}

func localTraceProvider() fx.Option {
	return fx.Provide(
		sdktrace.NeverSample,
		xtrace.NewStdoutTracingInitializer,
	)
}

type InjectParam struct {
	sqlc.DBConfig
}

func InjectLocal(param InjectParam, f func(e *rest.Extcuter)) *fx.App {
	return fx.New(
		fx.Supply(param.DBConfig),
		localRepositoryProvider(),
		localAuthServiceProvider(),
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
