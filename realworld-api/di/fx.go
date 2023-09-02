package di

import (
	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	articlerepo "github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authrepo "github.com/ryutah/realworld-echo/realworld-api/domain/auth/repository"
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

func localAuthRepositoryProvider() fx.Option {
	return fx.Provide(
		fx.Annotate(firebase.NewUser, fx.As(new(authrepo.User))),
	)
}

func localArticleRepositoryProvider() fx.Option {
	return fx.Provide(
		fx.Annotate(sqlc.NewDBManager, fx.As(new(sqlc.DBManager))),
		fx.Annotate(sqlc.NewArtile, fx.As(new(articlerepo.Article))),
		fx.Annotate(sqlc.NewFavorite, fx.As(new(articlerepo.Favorite))),
		fx.Annotate(sqlc.NewFollow, fx.As(new(articlerepo.Follow))),
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
		localArticleRepositoryProvider(),
		localAuthServiceProvider(),
		localAuthRepositoryProvider(),
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
