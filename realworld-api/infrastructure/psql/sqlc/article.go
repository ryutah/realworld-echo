package sqlc

import (
	"context"

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
	repository  struct {
		user authrepo.User
	}
}

var _ repository.Article = (*Article)(nil)

func NewArtile(manager DBManager, userRepo authrepo.User) *Article {
	return &Article{
		manager: manager,
		repository: struct {
			user authrepo.User
		}{
			user: userRepo,
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

func (a *Article) Search(_ context.Context, _ repository.ArticleSearchParam) (model.ArticleSlice, error) {
	panic("not implemented") // TODO: Implement
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
