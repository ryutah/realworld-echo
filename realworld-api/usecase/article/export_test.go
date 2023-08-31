package article

import (
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
)

func (l ListArticleParam) ToSearchParam() (*repository.ArticleSearchParam, error) {
	return l.toSearchParam()
}

func (c CreateArticleParam) ToDomain(slug model.Slug, user *authmodel.User) (*model.Article, []model.Tag, error) {
	return c.toDomain(slug, user)
}
