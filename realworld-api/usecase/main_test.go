//go:generate mockgen -package=mock -destination=mock/usecase_mock.go github.com/ryutah/realworld-echo/realworld-api/usecase GetArticleOutputPort,ErrorOutputPort,ErrorHandler,ErrorReporter
//go:generate mockgen -package=mock -destination=mock/repository_mock.go github.com/ryutah/realworld-echo/realworld-api/domain/repository Article

package usecase_test
