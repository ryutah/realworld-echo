package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type getArticleOutputPort struct{}

func NewGetArticleOutputPort(e usecase.ErrorOutputPort) usecase.OKOutputPort[usecase.GetArticleResult] {
	return &getArticleOutputPort{}
}

func (g *getArticleOutputPort) OK(ctx context.Context, _ usecase.GetArticleResult) error {
	c := echoContextFromContext(ctx)
	return c.JSON(http.StatusOK, gen.SingleArticleResponse{
		Article: gen.Article{
			Author: gen.Profile{
				Bio:       "dummy",
				Following: false,
				Image:     "dummy",
				Username:  "dummy",
			},
			Body:           "dummy",
			CreatedAt:      time.Now(),
			Description:    "dummy",
			Favorited:      false,
			FavoritesCount: 0,
			Slug:           "dummy",
			TagList:        []string{},
			Title:          "",
			UpdatedAt:      time.Now(),
		},
	})
}

type Article struct {
	inputPort struct {
		getArticle usecase.GetArticleInputPort
	}
}

func NewArticle(getArticle usecase.GetArticleInputPort) *Article {
	return &Article{
		inputPort: struct {
			getArticle usecase.GetArticleInputPort
		}{
			getArticle: getArticle,
		},
	}
}

func (a *Article) GetArticle(c echo.Context, slug string) error {
	ctx := newContext(c)
	ctx, span := xtrace.StartSpan(ctx)
	defer span.End()
	return a.inputPort.getArticle.Get(ctx, slug)
}
