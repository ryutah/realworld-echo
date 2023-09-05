package article

import (
	"context"
	"errors"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/internal/operations"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/samber/lo"
)

type (
	ListArticleParam struct {
		Tag         string
		Author      string
		FavoritedBy string
		Limit       uint
		Offset      uint
	}
	ListArticleResultArtile struct {
		Aritcle         model.Article
		FavoriteCount   int
		Favorited       bool
		AuthorFollowing bool
	}
	ListArticleResult struct {
		Articles []ListArticleResultArtile
	}
	ListArticleInputPort interface {
		List(context.Context, ListArticleParam) *usecase.Result[ListArticleResult]
	}
)

type ListArticle struct {
	errorHandler usecase.ErrorHandler[ListArticleResult]
	repository   struct {
		article  repository.Article
		favorite repository.Favorite
		follow   repository.Follow
	}
	service struct {
		auth service.Auth
	}
}

func NewListArticle(
	errorHandler usecase.ErrorHandler[ListArticleResult],
	articleRepo repository.Article,
	favoriteRepo repository.Favorite,
	followRepo repository.Follow,
	authService service.Auth,
) ListArticleInputPort {
	return &ListArticle{
		errorHandler: errorHandler,
		repository: struct {
			article  repository.Article
			favorite repository.Favorite
			follow   repository.Follow
		}{
			article:  articleRepo,
			favorite: favoriteRepo,
			follow:   followRepo,
		},
		service: struct {
			auth service.Auth
		}{
			auth: authService,
		},
	}
}

func (a *ListArticle) List(ctx context.Context, param ListArticleParam) *usecase.Result[ListArticleResult] {
	ctx, finish := operations.StartFunc(ctx, operations.FuncParam(param))
	defer finish()

	searchParam, err := param.toSearchParam()
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	articles, err := a.repository.article.Search(ctx, *searchParam)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	favoriteCounts, err := a.repository.favorite.CountList(ctx, articles.Slugs()...)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	user, err := a.service.auth.CurrentUser(ctx)
	if errors.Is(err, derrors.Errors.NotAuthorized.Err) {
		return usecase.Success(a.generateResult(articles, favoriteCounts, nil, nil, nil))
	} else if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}

	favorited, err := a.repository.favorite.ExistsList(ctx, user.ID, articles.Slugs()...)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	follows, err := a.repository.follow.ExistsList(ctx, user.ID, articles.Authors()...)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	return usecase.Success(a.generateResult(articles, favoriteCounts, favorited, follows, user))
}

func (a *ListArticle) generateResult(articles model.ArticleSlice, favoriteCnt model.FavoriteCountMap, favorited model.FavoriteExistsMap, follows model.FollowersExistsMap, user *authmodel.User) ListArticleResult {
	artileResults := lo.Map(articles, func(item model.Article, _ int) ListArticleResultArtile {
		var following bool
		if user != nil {
			following = follows.IsFollowing(item.Author.ID)
		}
		return ListArticleResultArtile{
			Aritcle:         item,
			FavoriteCount:   favoriteCnt.Count(item.Slug),
			Favorited:       favorited.Exists(item.Slug),
			AuthorFollowing: following,
		}
	})
	return ListArticleResult{
		Articles: artileResults,
	}
}

func (l ListArticleParam) toSearchParam() (*repository.ArticleSearchParam, error) {
	var (
		tag          *model.TagName
		pauthor      *authmodel.UserID
		pfavoritedBy *authmodel.UserID
		offset       premitive.Offset
		limit        = repository.DefaultLimit
		err          error
	)

	if l.Tag != "" {
		t, err := model.NewTagName(l.Tag)
		if err != nil {
			return nil, err
		}
		tag = &t
	}
	if l.Author != "" {
		author, err := authmodel.NewUserID(l.Author)
		if err != nil {
			return nil, err
		}
		pauthor = &author
	}
	if l.FavoritedBy != "" {
		favoritedBy, err := authmodel.NewUserID(l.FavoritedBy)
		if err != nil {
			return nil, err
		}
		pfavoritedBy = &favoritedBy
	}
	if l.Offset > 0 {
		offset = premitive.NewOffset(l.Offset)
	}
	if l.Limit > 0 {
		limit, err = premitive.NewLimit(l.Limit)
		if err != nil {
			return nil, err
		}
	}

	return &repository.ArticleSearchParam{
		Tag:         tag,
		Author:      pauthor,
		FavoritedBy: pfavoritedBy,
		Offset:      offset,
		Limit:       limit,
	}, nil
}
