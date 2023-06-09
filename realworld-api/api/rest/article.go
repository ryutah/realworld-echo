package rest

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/usecase/article"
)

type getArticleOutputPort struct {
	usecase.GenericsErrorOutputPort[article.GetArticleResult]
}

func NewGetArticleOutputPort(e usecase.ErrorOutputPort) usecase.OutputPort[article.GetArticleResult] {
	return &getArticleOutputPort{
		GenericsErrorOutputPort: newGenericsErrorOutputPort[article.GetArticleResult](),
	}
}

func (g *getArticleOutputPort) Success(ctx context.Context, article article.GetArticleResult) error {
	c := echoContextFromContext(ctx)
	return c.JSON(http.StatusOK, gen.SingleArticleResponse{
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
			CreatedAt:      article.Article.CreatedAt,
			UpdatedAt:      article.Article.UpdatedAt,
		},
	})
}

type Article struct {
	inputPort struct {
		getArticle  article.GetArticleInputPort
		listArticle article.List[gen.GetArticlesResponseObject]
	}
}

func NewArticle(getArticle article.GetArticleInputPort) *Article {
	return &Article{
		inputPort: struct {
			getArticle  article.GetArticleInputPort
			listArticle article.List[gen.GetArticlesResponseObject]
		}{
			getArticle: getArticle,
		},
	}
}

func (a *Article) GetArticles(ctx context.Context, request gen.GetArticlesRequestObject) (gen.GetArticlesResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Article) GetArticle_(c echo.Context, slug string) error {
	ctx, span := xtrace.StartSpan(newContext(c))
	defer span.End()
	return a.inputPort.getArticle.Get(ctx, slug)
}

func (a *Article) CreateArticle(c echo.Context) error {
	_, span := xtrace.StartSpan(newContext(c))
	defer span.End()

	var req gen.NewArticleRequest
	if ge, ok := bindAndValidateBody(c, &req); !ok {
		return c.JSON(echo.ErrBadRequest.Code, ge)
	}
	panic("not implemented") // TODO: Implement
}

func bindAndValidateBody(ctx echo.Context, v any) (gen.GenericError, bool) {
	if err := ctx.Bind(&v); err != nil {
		return gen.GenericError{
			Errors: struct {
				Body []string `json:"body"`
			}{
				Body: []string{err.Error()},
			},
		}, false
	}

	if err := validator.New().Struct(v); err != nil {
		ves := err.(validator.ValidationErrors)
		msgs := make([]string, len(ves))
		for _, ve := range ves {
			msgs = append(msgs, ve.Error())
		}
		return gen.GenericError{
			Errors: struct {
				Body []string "json:\"body\""
			}{
				Body: msgs,
			},
		}, false
	}
	return gen.GenericErrorModel{}, true
}
