package article

import "github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"

func (l ListArticleParam) ToSearchParam() (*repository.ArticleSearchParam, error) {
	return l.toSearchParam()
}
