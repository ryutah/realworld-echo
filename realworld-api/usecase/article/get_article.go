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
	"go.uber.org/zap"
)

type (
	GetArticleResult struct {
		Article         model.Article
		Favorited       bool
		Favorites       model.FavoriteSlice
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
	ctx, finish := operations.StartFunc(ctx, zap.String("slug", slugStr))
	defer finish()

	slug, err := model.NewSlug(slugStr)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithNotFoundHandler(derrors.Errors.NotFound.Err))
	}

	favorites, err := a.repository.favorites.ListBySlug(ctx, article.Slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	user, err := a.service.auth.CurrentUser(ctx)
	if errors.Is(err, derrors.Errors.NotAuthorized.Err) {
		return usecase.Success(GetArticleResult{
			Article:   *article,
			Favorited: false,
			Favorites: favorites,
		})
	} else if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	follows, err := a.repository.follow.ExistsList(ctx, user.ID, article.Author)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	return usecase.Success(GetArticleResult{
		Article:         *article,
		Favorited:       favorites.IsFavorited(user.ID, article.Slug),
		Favorites:       favorites,
		FollowingAuthor: follows.IsFollowing(article.Author),
	})
}
