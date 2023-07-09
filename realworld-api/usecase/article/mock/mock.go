//go:generate mockgen -package=gen -destination=gen/usecase_article_mock.gen.go . GetArticleOutputPort,CreateArticleOutputPort,ListArticleOutputPort
//go:generate mockgen -package=gen -destination=gen/usecase_mock.gen.go github.com/ryutah/realworld-echo/realworld-api/usecase ErrorOutputPort,ErrorHandler,ErrorReporter
//go:generate mockgen -package=gen -destination=gen/article_repository_mock.gen.go github.com/ryutah/realworld-echo/realworld-api/domain/article/repository Article

package mock

import (
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article"
)

// NOTE: workfaround for gomock
// see: https://github.com/golang/mock/issues/621#issuecomment-1094351718
type (
	GetArticleOutputPort    = usecase.OutputPort[article.GetArticleResult]
	CreateArticleOutputPort = usecase.OutputPort[article.CreateArticleResult]
	ListArticleOutputPort   = usecase.NewOutputPort[article.ListResult, any]
)
