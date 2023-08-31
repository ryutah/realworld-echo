package sqlc

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc/gen"
)

type Article struct {
	manager DBManager
}

var _ repository.Article = (*Article)(nil)

func NewArtile(manager DBManager) *Article {
	return &Article{
		manager: manager,
	}
}

func (a *Article) GenerateID(_ context.Context) (model.Slug, error) {
	return model.NewSlug(uuid.New().String())
}

func (a *Article) Get(ctx context.Context, slug model.Slug) (*model.Article, error) {
	uuid, err := uuid.Parse(slug.String())
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to parse slug to uuid")
	}

	q := a.manager.Querier(ctx)

	article, err := q.GetArticle(ctx, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, derrors.NewNotFoundError(0, err, "failed to get article")
	} else if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get article")
	}

	articleTags, err := q.ListArticleTags(ctx, []string{article.Slug.String()})
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get article_tags")
	}

	return a.reCreateEntity(article, articleTags)
}

func (a *Article) Save(ctx context.Context, article model.Article) error {
	panic("not implemented") // TODO: Implement
}

func (a *Article) Search(_ context.Context, _ repository.ArticleSearchParam) (model.ArticleSlice, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Article) reCreateEntity(article gen.Article, tags []gen.ListArticleTagsRow) (*model.Article, error) {
	slug, err := model.NewSlug(article.Slug.String())
	if err != nil {
		return nil, err
	}
	author, err := authmodel.NewUserID(article.Author)
	if err != nil {
		return nil, err
	}
	contents, err := model.NewArticleContents(article.Title, article.Description, article.Body)
	if err != nil {
		return nil, err
	}
	var tagNames []model.TagName
	for _, t := range tags {
		tn, err := model.NewTagName(t.Name)
		if err != nil {
			return nil, err
		}
		tagNames = append(tagNames, tn)
	}
	return model.ReCreateArticle(
		slug,
		*contents,
		author,
		tagNames,
		premitive.NewJSTTime(article.CreatedAt),
		premitive.NewJSTTime(article.UpdatedAt),
	)
}
