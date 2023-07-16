package rest

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article"
)

type Article struct {
	inputPort struct {
		getArticle  article.GetArticleInputPort
		listArticle article.ListArticle[gen.GetArticlesResponseObject]
	}
}

func NewArticle(getArticle article.GetArticleInputPort) *Article {
	return &Article{
		inputPort: struct {
			getArticle  article.GetArticleInputPort
			listArticle article.ListArticle[gen.GetArticlesResponseObject]
		}{
			getArticle: getArticle,
		},
	}
}

func (a *Article) GetArticles(ctx context.Context, request gen.GetArticlesRequestObject) (gen.GetArticlesResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Article) GetArticle(ctx context.Context, request gen.GetArticleRequestObject) (gen.GetArticleResponseObject, error) {
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()
	result := a.inputPort.getArticle.Get(ctx, request.Slug)

	if !result.IsFailed() {
		article := result.Success()
		return gen.GetArticle200JSONResponse{
			SingleArticleResponseJSONResponse: gen.SingleArticleResponseJSONResponse{
				Article: gen.Article{
					Author: gen.Profile{
						Bio:       "dummy",
						Following: false,
						Image:     "dummy",
						Username:  "dummy",
					},
					Slug:           article.Article.Slug.String(),
					Title:          article.Article.Contents.Title.String(),
					Description:    article.Article.Contents.Description.String(),
					Body:           article.Article.Contents.Description.String(),
					Favorited:      false,
					FavoritesCount: 0,
					TagList:        []string{},
					CreatedAt:      article.Article.CreatedAt.Time(),
					UpdatedAt:      article.Article.UpdatedAt.Time(),
				},
			},
		}, nil
	}
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
