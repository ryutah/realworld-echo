package article

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/auth"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	GetAuthArticleResult    struct{}
	GetAuthArticleInputPort interface {
		Get(ctx context.Context, token auth.AuthToken, slug string) error
	}
)

type GetAuthArticleOutputPort = usecase.OutputPort[GetAuthArticleResult]

type AuthArticle struct {
	outputPort struct {
		get GetAuthArticleInputPort
	}
	repository struct {
		article repository.Article
	}
}
