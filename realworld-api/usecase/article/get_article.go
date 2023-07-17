package article

import (
	"context"
	"fmt"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"go.uber.org/zap"
)

type (
	GetArticleResult struct {
		Article   model.Article
		Favorites model.FavoriteSlice
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
	}
}

func NewGetArticle(
	errorHandler usecase.ErrorHandler[GetArticleResult],
	articleRepo repository.Article,
	favoriteRepo repository.Favorite,
) GetArticleInputPort {
	return &GetArticle{
		errorHandler: errorHandler,
		repository: struct {
			article   repository.Article
			favorites repository.Favorite
		}{
			article:   articleRepo,
			favorites: favoriteRepo,
		},
	}
}

func (a *GetArticle) Get(ctx context.Context, slugStr string) *usecase.Result[GetArticleResult] {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	ctx = xlog.ContextWithLogFields(ctx, zap.String("slug", slugStr))

	slug, err := model.NewSlug(slugStr)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	xlog.Info(ctx, fmt.Sprintf("get article by: %v", slug))
	article, err := a.repository.article.Get(ctx, slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithNotFoundHandler(derrors.Errors.NotFound.Err))
	}

	favorites, err := a.repository.favorites.ListBySlug(ctx, article.Slug)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	return usecase.Success(GetArticleResult{
		Article:   *article,
		Favorites: favorites,
	})
}
