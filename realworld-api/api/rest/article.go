package rest

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/internal/operations"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	"github.com/samber/lo"
)

type Article struct {
	inputPort struct {
		getArticle  article.GetArticleInputPort
		listArticle article.ListArticleInputPort
	}
}

func NewArticle(getArticle article.GetArticleInputPort, listArticle article.ListArticleInputPort) *Article {
	return &Article{
		inputPort: struct {
			getArticle  article.GetArticleInputPort
			listArticle article.ListArticleInputPort
		}{
			getArticle:  getArticle,
			listArticle: listArticle,
		},
	}
}

func (a *Article) GetArticles(ctx context.Context, request gen.GetArticlesRequestObject) (gen.GetArticlesResponseObject, error) {
	ctx, finish := operations.StartFunc(ctx)
	defer finish()

	result := a.inputPort.listArticle.List(ctx, article.ListArticleParam{})
	if result.IsFailed() {
		return gen.GetArticles422JSONResponse{
			GenericErrorJSONResponse: gen.GenericErrorJSONResponse{
				Errors: struct {
					Body []string `json:"body"`
				}{
					Body: []string{
						result.Fail().Message,
					},
				},
			},
		}, nil
	}
	return gen.GetArticles200JSONResponse{
		MultipleArticlesResponseJSONResponse: gen.MultipleArticlesResponseJSONResponse{
			Articles: lo.Map(result.Success().Articles, func(a article.ListArticleResultArtile, _ int) gen.Article {
				return gen.Article{
					Slug: a.Aritcle.Slug.String(),
					Author: gen.Profile{
						Bio:       a.Aritcle.Author.Bio.String(),
						Following: a.AuthorFollowing,
						Image:     a.Aritcle.Author.Image.String(),
						Username:  a.Aritcle.Author.Name.String(),
					},
					Title:          a.Aritcle.Contents.Title.String(),
					Body:           a.Aritcle.Contents.Body.String(),
					Description:    a.Aritcle.Slug.String(),
					Favorited:      a.Favorited,
					FavoritesCount: a.FavoriteCount,
					TagList:        lo.Map(a.Aritcle.Tags, func(t model.TagName, _ int) string { return t.String() }),
					CreatedAt:      a.Aritcle.CreatedAt.Time(),
					UpdatedAt:      a.Aritcle.UpdatedAt.Time(),
				}
			}),
			ArticlesCount: len(result.Success().Articles),
		},
	}, nil
}

func (a *Article) GetArticle(ctx context.Context, request gen.GetArticleRequestObject) (gen.GetArticleResponseObject, error) {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()
	result := a.inputPort.getArticle.Get(ctx, request.Slug)

	if result.IsFailed() {
		return gen.GetArticle422JSONResponse{
			GenericErrorJSONResponse: gen.GenericErrorJSONResponse{
				Errors: struct {
					Body []string `json:"body"`
				}{
					Body: []string{
						result.Fail().Message,
					},
				},
			},
		}, nil
	}

	article := result.Success()
	return gen.GetArticle200JSONResponse{
		SingleArticleResponseJSONResponse: gen.SingleArticleResponseJSONResponse{
			Article: gen.Article{
				Author: gen.Profile{
					Bio:       article.Article.Author.Bio.String(),
					Following: article.FollowingAuthor,
					Image:     article.Article.Author.Image.String(),
					Username:  article.Article.Author.Name.String(),
				},
				Slug:           article.Article.Slug.String(),
				Title:          article.Article.Contents.Title.String(),
				Description:    article.Article.Contents.Description.String(),
				Body:           article.Article.Contents.Description.String(),
				Favorited:      article.Favorited,
				FavoritesCount: article.FavoriteCount,
				TagList:        lo.Map(article.Article.Tags, func(t model.TagName, _ int) string { return t.String() }),
				CreatedAt:      article.Article.CreatedAt.Time(),
				UpdatedAt:      article.Article.UpdatedAt.Time(),
			},
		},
	}, nil
}
