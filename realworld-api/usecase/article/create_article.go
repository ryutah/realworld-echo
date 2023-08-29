package article

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/auth/service"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
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
	repository   struct {
		article repository.Article
	}
	service struct {
		auth service.Auth
	}
}

func NewCreateArticle(
	errorHandler usecase.ErrorHandler[CreateArticleResult],
	articleRepo repository.Article,
	authService service.Auth,
) CreateArticleInputPort {
	return &CreateArticle{
		errorHandler: errorHandler,
		repository: struct {
			article repository.Article
		}{
			article: articleRepo,
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

	newArticle, err := param.toDomain(newSlug, user)
	if err != nil {
		return c.errorHandler.Handle(ctx, err, usecase.WithBadRequestHandler(derrors.Errors.Validation.Err))
	}

	if err := c.repository.article.Save(ctx, *newArticle); err != nil {
		return c.errorHandler.Handle(ctx, err)
	}
	return usecase.Success(CreateArticleResult{
		Article: *newArticle,
	})
}

func (c CreateArticleParam) toDomain(slug model.Slug, user *authmodel.User) (*model.Article, error) {
	var tags []model.ArticleTag
	for _, t := range c.Tags {
		tag, err := model.NewArticleTag(t)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}
	content, err := model.NewArticleContents(c.Title, c.Description, c.Body, tags)
	if err != nil {
		return nil, err
	}
	return model.NewArticle(slug, *content, user.ID)
}
