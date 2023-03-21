package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/api/rest/gen"
	"github.com/ryutah/realworld-echo/usecase"
)

type getArticleOutputPort struct {
	usecase.ErrorOutputPort
}

func NewGetArticleOutputPort(e usecase.ErrorOutputPort) usecase.GetArticleOutputPort {
	return &getArticleOutputPort{
		ErrorOutputPort: e,
	}
}

func (g *getArticleOutputPort) Ok(ctx context.Context, _ usecase.GetArticleResult) error {
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

type Artcle struct {
	inputPort struct {
		getArticle usecase.GetArticleInputPort
	}
}

func NewArticle(getArticle usecase.GetArticleInputPort) *Artcle {
	return &Artcle{
		inputPort: struct {
			getArticle usecase.GetArticleInputPort
		}{
			getArticle: getArticle,
		},
	}
}

func (a *Artcle) GetArticle(c echo.Context, slug string) error {
	ctx := newContext(c)
	return a.inputPort.getArticle.Get(ctx, slug)
}
