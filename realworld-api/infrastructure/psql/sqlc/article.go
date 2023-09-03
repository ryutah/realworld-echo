package sqlc

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	authrepo "github.com/ryutah/realworld-echo/realworld-api/domain/auth/repository"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/domain/transaction"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc/gen"
	"github.com/samber/lo"
)

type Article struct {
	transaction transaction.Transaction
	manager     DBManager
	selector    struct {
		articles RawSelector[gen.Article]
	}
	repository struct {
		user authrepo.User
	}
}

var _ repository.Article = (*Article)(nil)

func NewArtile(
	manager DBManager,
	transaction transaction.Transaction,
	userRepo authrepo.User,
	articleSelector RawSelector[gen.Article],
) *Article {
	return &Article{
		manager:     manager,
		transaction: transaction,
		repository: struct {
			user authrepo.User
		}{
			user: userRepo,
		},
		selector: struct {
			articles RawSelector[gen.Article]
		}{
			articles: articleSelector,
		},
	}
}

func (a *Article) GenerateID(_ context.Context) (model.Slug, error) {
	return model.NewSlug(uuid.New().String())
}

func (a *Article) Get(ctx context.Context, slug model.Slug) (*model.Article, error) {
	uid, err := uuid.Parse(slug.String())
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to parse slug to uuid")
	}

	q := a.manager.Querier(ctx)

	article, err := q.GetArticle(ctx, uid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, derrors.NewNotFoundError(0, err, "failed to get article")
	} else if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get article")
	}

	articleTags, err := q.ListArticleTags(ctx, []uuid.UUID{article.Slug})
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get article_tags")
	}

	author, err := a.repository.user.Get(ctx, authmodel.UserID(article.Author))
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get author")
	}

	return a.reCreateEntity(article, *author, articleTags)
}

func (a *Article) Save(ctx context.Context, article model.Article) error {
	slug, err := uuid.Parse(article.Slug.String())
	if err != nil {
		return derrors.NewInternalError(0, err, "failed to parse slug to uuid")
	}

	if err := a.transaction.Run(ctx, func(tc context.Context) error {
		q := a.manager.Querier(ctx)

		if err := q.UpsertArticle(ctx, gen.UpsertArticleParams{
			Slug:        slug,
			Author:      article.Author.ID.String(),
			Body:        article.Contents.Body.String(),
			Title:       article.Contents.Title.String(),
			Description: article.Contents.Description.String(),
			CreatedAt:   toTimestamptz(article.CreatedAt.Time()),
			UpdatedAt:   toTimestamptz(article.UpdatedAt.Time()),
		}); err != nil {
			return derrors.NewInternalError(0, err, "failed to upsert article")
		}
		if err := q.DeleteArticleTagBySlug(ctx, slug); err != nil {
			return derrors.NewInternalError(0, err, "failed to delete article_tags")
		}
		params := lo.Map(article.Tags, func(item model.TagName, index int) gen.CreateArticleTagParams {
			return gen.CreateArticleTagParams{
				ArticleSlug: slug,
				TagName:     item.String(),
				CreatedAt:   toTimestamptz(article.CreatedAt.Time()),
				UpdatedAt:   toTimestamptz(article.UpdatedAt.Time()),
			}
		})
		if _, err := q.CreateArticleTag(ctx, params); err != nil {
			return derrors.NewInternalError(0, err, "failed to create article_tags")
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

type articleSearchParam repository.ArticleSearchParam

func (a articleSearchParam) condition() squirrel.Eq {
	eqCond := make(squirrel.Eq)
	if a.Author != nil {
		eqCond["a.author"] = a.Author.String()
	}
	if a.FavoritedBy != nil {
		eqCond["f.user_id"] = a.FavoritedBy.String()
	}
	if a.Tag != nil {
		eqCond["t.tag_name"] = a.Tag.String()
	}
	return eqCond
}

func (a *Article) Search(ctx context.Context, param repository.ArticleSearchParam) (model.ArticleSlice, error) {
	p := articleSearchParam(param)

	articles, err := a.selector.articles.Select(
		ctx, a.manager.Executor(ctx),
		squirrel.Select("a.*").
			From("article as a").
			LeftJoin("article_favorite as f on a.slug = f.article_slug").
			LeftJoin("article_tag as t on a.slug = t.article_slug").
			Where(p.condition()).
			Limit(uint64(param.Limit)).
			Offset(uint64(param.Offset)),
	)
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to select articles")
	}

	articleTags, err := a.manager.Querier(ctx).ListArticleTags(ctx, lo.Map(articles, func(i gen.Article, _ int) uuid.UUID {
		return i.Slug
	}))
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get article_tags")
	}
	authors, err := a.repository.user.List(ctx, lo.Map(articles, func(i gen.Article, _ int) authmodel.UserID {
		return authmodel.UserID(i.Author)
	})...)
	if err != nil {
		return nil, derrors.NewInternalError(0, err, "failed to get authors")
	}

	return a.reCreateEntities(articles, authors, articleTags)
}

func (a *Article) reCreateEntities(articles []gen.Article, authors []authmodel.User, tags []gen.ListArticleTagsRow) ([]model.Article, error) {
	var results []model.Article
	for _, article := range articles {
		author, _ := lo.Find(authors, func(i authmodel.User) bool {
			return article.Author == i.ID.String()
		})
		tag := lo.Filter(tags, func(i gen.ListArticleTagsRow, _ int) bool {
			return article.Slug == i.ArticleSlug
		})
		result, err := a.reCreateEntity(article, author, tag)
		if err != nil {
			return nil, err
		}
		results = append(results, *result)
	}
	return results, nil
}

func (a *Article) reCreateEntity(article gen.Article, author authmodel.User, tags []gen.ListArticleTagsRow) (*model.Article, error) {
	slug, err := model.NewSlug(article.Slug.String())
	if err != nil {
		return nil, err
	}
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
		*model.NewUserProfile(author),
		tagNames,
		premitive.NewJSTTime(article.CreatedAt.Time),
		premitive.NewJSTTime(article.UpdatedAt.Time),
	)
}
