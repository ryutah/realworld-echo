package article_test

import "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"

func mustNewTag(s string) *model.ArticleTag {
	tag, err := model.NewArticleTag(s)
	if err != nil {
		panic(err)
	}
	return tag
}
