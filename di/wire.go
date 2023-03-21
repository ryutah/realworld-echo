//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ryutah/realworld-echo/api/rest"
	"github.com/ryutah/realworld-echo/usecase"
)

var (
	restSet = wire.NewSet(
		rest.NewServer,
		rest.NewArticle,
		inputPortSet,
	)
	outputPortSet = wire.NewSet(
		rest.NewErrorOutputPort,
		rest.NewGetArticleOutputPort,
	)
	inputPortSet = wire.NewSet(
		usecase.NewArticle,
		wire.Bind(new(usecase.GetArticleInputPort), new(*usecase.Article)),
		outputPortSet,
	)
)

func InitializeServer() *rest.Server {
	panic(wire.Build(restSet))
}
