package article

import (
	"context"
	"errors"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/internal/operations"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	GetArticleResult struct {
		Article         model.Article
		Favorited       bool
		FavoriteCount   int
		FollowingAuthor bool
	}
	GetArticleInputPort interface {
		Get(ctx context.Context, slugStr string) *usecase.Result[GetArticleResult]
	}
)

type GetArticle struct {
	errorHandler usecase.ErrorHandler[GetArticleResult]
	repository   struct {
		article   repository.Article
		favorites repository.Favorite
		follow    repository.Follow
	}
	service struct {
		auth service.Auth
	}
}

func NewGetArticle(
	errorHandler usecase.ErrorHandler[GetArticleResult],
	articleRepo repository.Article,
	favoriteRepo repository.Favorite,
	followRepo repository.Follow,
	authService service.Auth,
) GetArticleInputPort {
	return &GetArticle{
		errorHandler: errorHandler,
		repository: struct {
			article   repository.Article
			favorites repository.Favorite
			follow    repository.Follow
		}{
			article:   articleRepo,
			favorites: favoriteRepo,
			follow:    followRepo,
		},
		service: struct {
			auth service.Auth
		}{
			auth: authService,
		},
	}
}

func (a *GetArticle) Get(ctx context.Context, slugStr string) *usecase.Result[GetArticleResult] {
	ctx, finish := operations.StartFunc(ctx, operations.FuncParam(slugStr))
	defer finish()

	slug, err := model.NewSlug(slugStr)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithNotFoundHandler(derrors.Errors.NotFound.Err))
	}

	favoriteCount, err := a.repository.favorites.Count(ctx, article.Slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	user, err := a.service.auth.CurrentUser(ctx)
	if errors.Is(err, derrors.Errors.NotAuthorized.Err) {
		return usecase.Success(GetArticleResult{
			Article:       *article,
			Favorited:     false,
			FavoriteCount: favoriteCount,
		})
	} else if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	favorited, err := a.repository.favorites.Exists(ctx, user.ID, article.Slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	following, err := a.repository.follow.Exists(ctx, user.ID, article.Author.ID)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	return usecase.Success(GetArticleResult{
		Article:         *article,
		Favorited:       favorited,
		FavoriteCount:   favoriteCount,
		FollowingAuthor: following,
	})
}
