package article

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/transaction"
	"github.com/ryutah/realworld-echo/realworld-api/internal/operations"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
)

type (
	CreateArticleInputPort interface {
		Create(context.Context, CreateArticleParam) *usecase.Result[CreateArticleResult]
	}
	CreateArticleParam struct {
		Title       string
		Description string
		Body        string
		Tags        []string
	}
	CreateArticleResult struct {
		Article model.Article
	}
)

type CreateArticle struct {
	errorHandler usecase.ErrorHandler[CreateArticleResult]
	transaction  transaction.Transaction
	repository   struct {
		article repository.Article
		tag     repository.Tag
	}
	service struct {
		auth service.Auth
	}
}

func NewCreateArticle(
	errorHandler usecase.ErrorHandler[CreateArticleResult],
	transaction transaction.Transaction,
	articleRepo repository.Article,
	tagRepo repository.Tag,
	authService service.Auth,
) CreateArticleInputPort {
	return &CreateArticle{
		errorHandler: errorHandler,
		transaction:  transaction,
		repository: struct {
			article repository.Article
			tag     repository.Tag
		}{
			article: articleRepo,
			tag:     tagRepo,
		},
		service: struct {
			auth service.Auth
		}{
			auth: authService,
		},
	}
}

func (c *CreateArticle) Create(ctx context.Context, param CreateArticleParam) *usecase.Result[CreateArticleResult] {
	ctx, finish := operations.StartFunc(ctx)
	defer finish()

	user, err := c.service.auth.CurrentUser(ctx)
	if err != nil {
		return c.errorHandler.Handle(ctx, err, usecase.WithUnauthorizedHandler(derrors.Errors.NotAuthorized.Err))
	}
	newSlug, err := c.repository.article.GenerateID(ctx)
	if err != nil {
		return c.errorHandler.Handle(ctx, err)
	}
	newArticle, tags, err := param.toDomain(newSlug, user)
	if err != nil {
		return c.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	if err := c.transaction.Run(ctx, func(tc context.Context) error {
		if err := c.repository.article.Save(tc, *newArticle); err != nil {
			return err
		}
		return c.repository.tag.BulkSave(ctx, tags)
	}); err != nil {
		return c.errorHandler.Handle(ctx, err)
	}

	return usecase.Success(CreateArticleResult{
		Article: *newArticle,
	})
}

func (c CreateArticleParam) toDomain(slug model.Slug, user *authmodel.User) (*model.Article, []model.Tag, error) {
	var (
		names []model.TagName
		tags  []model.Tag
	)
	for _, t := range c.Tags {
		name, nerr := model.NewTagName(t)
		names = append(names, name)
		tag, terr := model.NewTag(name)
		if nerr != nil || terr != nil {
			return nil, nil, errors.CombineErrors(nerr, terr)
		}
		tags = append(tags, *tag)
	}
	content, err := model.NewArticleContents(c.Title, c.Description, c.Body)
	if err != nil {
		return nil, nil, err
	}
	article, err := model.NewArticle(slug, *content, user.ID, names)
	if err != nil {
		return nil, nil, err
	}
	return article, tags, nil
}
