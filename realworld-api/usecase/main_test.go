//go:generate mockgen -package=mock -destination=mock/usecase_mock.go github.com/ryutah/realworld-echo/realworld-api/usecase ErrorOutputPort,ErrorHandler,ErrorReporter

package usecase_test
