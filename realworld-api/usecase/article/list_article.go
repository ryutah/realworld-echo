package article

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	ListArticleParam struct {
		Tag         string
		Author      string
		FavoritedBy string
	}
	ListArticleResult struct {
		Articles []model.Article
	}
	ListArticleInputPort interface {
		List(context.Context, ListArticleParam) *usecase.Result[ListArticleResult]
	}
)

func (l ListArticleParam) toSearchParam() (*repository.ArticleSearchParam, error) {
	var (
		tag          *model.ArticleTag
		pauthor      *authmodel.UserID
		pfavoritedBy *authmodel.UserID
		err          error
	)

	if l.Tag != "" {
		tag, err = model.NewArticleTag(l.Tag)
		if err != nil {
			return nil, err
		}
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

	return &repository.ArticleSearchParam{
		Tag:         tag,
		Author:      pauthor,
		FavoritedBy: pfavoritedBy,
	}, nil
}

type ListArticle[Ret any] struct {
	errorHandler usecase.ErrorHandler[ListArticleResult]
	repository   struct {
		article repository.Article
	}
}

func NewListArticle[Ret any](errorHandler usecase.ErrorHandler[ListArticleResult], articleRepo repository.Article) ListArticleInputPort {
	return &ListArticle[Ret]{
		errorHandler: errorHandler,
		repository: struct {
			article repository.Article
		}{
			article: articleRepo,
		},
	}
}

func (a *ListArticle[Ret]) List(ctx context.Context, param ListArticleParam) *usecase.Result[ListArticleResult] {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()

	searchParam, err := param.toSearchParam()
	if err != nil {
		return a.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	articles, err := a.repository.article.Search(ctx, *searchParam)
	if err != nil {
		return a.errorHandler.Handle(ctx, err)
	}
	// TODO(ryutah): お気に入りリストを取得する

	return usecase.Success(ListArticleResult{
		Articles: articles,
	})
}
