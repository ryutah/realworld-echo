package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/api/rest/gen"
)

type Server struct {
	gen.ServerInterface
	article *Article
}

func NewServer(article *Article) *Server {
	return &Server{
		article: article,
	}
}

func (a *Server) GetArticle(ctx echo.Context, slug string) error {
	return a.article.GetArticle(ctx, slug)
}
