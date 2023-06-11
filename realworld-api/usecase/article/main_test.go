//go:generate mockgen -package=mock -destination=mock/usecase_mock.go github.com/ryutah/realworld-echo/realworld-api/usecase ErrorOutputPort,ErrorHandler,ErrorReporter
//go:generate mockgen -package=mock -destination=mock/usecase_article_mock.go github.com/ryutah/realworld-echo/realworld-api/usecase/article GetArticleOutputPort
//go:generate mockgen -package=mock -destination=mock/article_repository_mock.go github.com/ryutah/realworld-echo/realworld-api/domain/article/repository Article

package article_test
