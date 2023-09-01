package article_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	mock_article_repo "github.com/ryutah/realworld-echo/realworld-api/internal/mock/article/repository"
	mock_auth_service "github.com/ryutah/realworld-echo/realworld-api/internal/mock/auth/service"
	mock_transaction "github.com/ryutah/realworld-echo/realworld-api/internal/mock/transaction"
	mock_usecase "github.com/ryutah/realworld-echo/realworld-api/internal/mock/usecase"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateArticle_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		param CreateArticleParam
	}
	type mock_authService struct {
		currentUser_returns_user  *authmodel.User
		currentUser_returns_error error
	}
	type mock_articleRepository struct {
		generateID_returns_slug  model.Slug
		generateID_returns_error error
		save_args_article        model.Article
		save_returns_error       error
	}
	type mock_tagRepository struct {
		bulkSave_args_tags     []model.Tag
		bulkSave_returns_error error
	}
	type mock_errorHandler struct {
		handle_args_error       error
		handle_args_opts_length int
		handle_returns_result   *usecase.Result[CreateArticleResult]
	}
	type mocks struct {
		articleRepository mock_articleRepository
		tagRepository     mock_tagRepository
		authService       mock_authService
		errorHandler      mock_errorHandler
	}
	type opts struct {
		nowFunc func() time.Time
	}
	type configs struct {
		errorHandler_handle_should_call      bool
		article_generateID_should_be_skipped bool
		article_save_should_be_skipped       bool
		tag_bulkSave_should_be_skipped       bool
		transaction_run_should_be_skipped    bool
	}

	var (
		now                     = time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
		nowJST                  = premitive.NewJSTTime(now)
		nowFunc                 = func() time.Time { return now }
		badRequestFailResult    = usecase.Fail[CreateArticleResult](usecase.NewFailResult(usecase.FailTypeBadRequest, "badrequest"))
		unauthorizedFailResult  = usecase.Fail[CreateArticleResult](usecase.NewFailResult(usecase.FailTypeUnauthorized, "unauthorized"))
		dummyError              = errors.New("error")
		internalErrorFailResult = usecase.Fail[CreateArticleResult](usecase.NewFailResult(usecase.FailTypeInternalError, "error"))
		testData1               = struct {
			args  args
			mocks mocks
			want  *usecase.Result[CreateArticleResult]
		}{
			args: args{
				param: CreateArticleParam{
					Title:       "title",
					Description: "desc",
					Body:        "body",
					Tags: []string{
						"tag1", "tag2",
					},
				},
			},
			mocks: mocks{
				authService: mock_authService{
					currentUser_returns_user: &authmodel.User{
						ID: "user1",
					},
				},
				articleRepository: mock_articleRepository{
					generateID_returns_slug: "slug",
					save_args_article: model.Article{
						Slug:   "slug",
						Author: "user1",
						Contents: model.ArticleContents{
							Title:       "title",
							Description: "desc",
							Body:        "body",
						},
						Tags: []model.TagName{
							"tag1", "tag2",
						},
						CreatedAt: nowJST,
						UpdatedAt: nowJST,
					},
				},
				tagRepository: mock_tagRepository{
					bulkSave_args_tags: []model.Tag{
						{
							Tag:       "tag1",
							CreatedAt: nowJST,
							UpdatedAt: nowJST,
						},
						{
							Tag:       "tag2",
							CreatedAt: nowJST,
							UpdatedAt: nowJST,
						},
					},
				},
			},
			want: usecase.Success(CreateArticleResult{
				Article: model.Article{
					Slug:   "slug",
					Author: "user1",
					Contents: model.ArticleContents{
						Title:       "title",
						Description: "desc",
						Body:        "body",
					},
					Tags: []model.TagName{
						"tag1", "tag2",
					},
					CreatedAt: nowJST,
					UpdatedAt: nowJST,
				},
			}),
		}
	)

	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    *usecase.Result[CreateArticleResult]
		opts    opts
		configs configs
	}{
		{
			name:  "when_given_valid_params_should_return_expect_success_result",
			args:  testData1.args,
			mocks: testData1.mocks,
			want:  testData1.want,
			opts: opts{
				nowFunc: nowFunc,
			},
		},
		{
			name: "when_given_invalid_params_should_call_erroHandler_handler_and_return_validation_error_result",
			args: args{
				param: CreateArticleParam{
					Title:       strings.Repeat("a", 10000),
					Description: testData1.args.param.Description,
					Body:        testData1.args.param.Body,
					Tags:        testData1.args.param.Tags,
				},
			},
			mocks: mocks{
				authService: testData1.mocks.authService,
				articleRepository: mock_articleRepository{
					generateID_returns_slug: testData1.mocks.articleRepository.generateID_returns_slug,
				},
				errorHandler: mock_errorHandler{
					handle_args_error:       derrors.Errors.Validation.Err,
					handle_args_opts_length: 1,
					handle_returns_result:   badRequestFailResult,
				},
			},
			want: badRequestFailResult,
			opts: opts{
				nowFunc: nowFunc,
			},
			configs: configs{
				transaction_run_should_be_skipped: true,
				article_save_should_be_skipped:    true,
				tag_bulkSave_should_be_skipped:    true,
				errorHandler_handle_should_call:   true,
			},
		},
		{
			name: "when_authService_currentUser_return_unauthorized_error_should_call_errorHandler_handler_and_return_unauthorized_fail_result",
			args: testData1.args,
			mocks: mocks{
				authService: mock_authService{
					currentUser_returns_error: derrors.Errors.NotAuthorized.Err,
				},
				errorHandler: mock_errorHandler{
					handle_args_error:       derrors.Errors.NotAuthorized.Err,
					handle_args_opts_length: 1,
					handle_returns_result:   unauthorizedFailResult,
				},
			},
			want: unauthorizedFailResult,
			opts: opts{
				nowFunc: nowFunc,
			},
			configs: configs{
				article_generateID_should_be_skipped: true,
				transaction_run_should_be_skipped:    true,
				article_save_should_be_skipped:       true,
				tag_bulkSave_should_be_skipped:       true,
				errorHandler_handle_should_call:      true,
			},
		},
		{
			name: "when_authService_currentUser_return_unknown_error_should_call_errorHandler_handler_and_return_internal_fail_result",
			args: testData1.args,
			mocks: mocks{
				authService: mock_authService{
					currentUser_returns_error: dummyError,
				},
				errorHandler: mock_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 1,
					handle_returns_result:   internalErrorFailResult,
				},
			},
			want: internalErrorFailResult,
			opts: opts{
				nowFunc: nowFunc,
			},
			configs: configs{
				article_generateID_should_be_skipped: true,
				transaction_run_should_be_skipped:    true,
				article_save_should_be_skipped:       true,
				tag_bulkSave_should_be_skipped:       true,
				errorHandler_handle_should_call:      true,
			},
		},
		{
			name: "when_artileRepository_generateID_return_unknown_error_should_call_errorHandler_handler_and_return_internal_fail_result",
			args: testData1.args,
			mocks: mocks{
				authService: testData1.mocks.authService,
				articleRepository: mock_articleRepository{
					generateID_returns_error: dummyError,
				},
				errorHandler: mock_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorFailResult,
				},
			},
			want: internalErrorFailResult,
			opts: opts{
				nowFunc: nowFunc,
			},
			configs: configs{
				transaction_run_should_be_skipped: true,
				article_save_should_be_skipped:    true,
				tag_bulkSave_should_be_skipped:    true,
				errorHandler_handle_should_call:   true,
			},
		},
		{
			name: "when_tagRepository_bulkSave_return_unknown_error_should_call_errorHandler_handler_and_return_internal_fail_result",
			args: testData1.args,
			mocks: mocks{
				authService: testData1.mocks.authService,
				articleRepository: mock_articleRepository{
					generateID_returns_slug: testData1.mocks.articleRepository.generateID_returns_slug,
				},
				tagRepository: mock_tagRepository{
					bulkSave_args_tags:     testData1.mocks.tagRepository.bulkSave_args_tags,
					bulkSave_returns_error: dummyError,
				},
				errorHandler: mock_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorFailResult,
				},
			},
			want: internalErrorFailResult,
			opts: opts{
				nowFunc: nowFunc,
			},
			configs: configs{
				article_save_should_be_skipped:  true,
				errorHandler_handle_should_call: true,
			},
		},
		{
			name: "when_artileRepository_save_return_unknown_error_should_call_errorHandler_handler_and_return_internal_fail_result",
			args: testData1.args,
			mocks: mocks{
				authService:   testData1.mocks.authService,
				tagRepository: testData1.mocks.tagRepository,
				articleRepository: mock_articleRepository{
					generateID_returns_slug: testData1.mocks.articleRepository.generateID_returns_slug,
					save_args_article:       testData1.mocks.articleRepository.save_args_article,
					save_returns_error:      dummyError,
				},
				errorHandler: mock_errorHandler{
					handle_args_error:       dummyError,
					handle_args_opts_length: 0,
					handle_returns_result:   internalErrorFailResult,
				},
			},
			want: internalErrorFailResult,
			opts: opts{
				nowFunc: nowFunc,
			},
			configs: configs{
				errorHandler_handle_should_call: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset := xtime.SetNowFunc(tt.opts.nowFunc)
			defer reset()

			authServiec := mock_auth_service.NewMockAuth(t)
			articleRepo := mock_article_repo.NewMockArticle(t)
			tagRepo := mock_article_repo.NewMockTag(t)
			transaction := mock_transaction.NewMockTransaction(t)
			errorHandler := mock_usecase.NewMockErrorHandler[CreateArticleResult](t)

			authServiec.EXPECT().
				CurrentUser(mock.Anything).
				Return(
					tt.mocks.authService.currentUser_returns_user,
					tt.mocks.authService.currentUser_returns_error,
				)
			if !tt.configs.article_generateID_should_be_skipped {
				articleRepo.EXPECT().
					GenerateID(mock.Anything).
					Return(
						tt.mocks.articleRepository.generateID_returns_slug,
						tt.mocks.articleRepository.generateID_returns_error,
					)
			}
			if !tt.configs.transaction_run_should_be_skipped {
				transactionExpectations(t, transaction)
			}
			if !tt.configs.article_save_should_be_skipped {
				articleRepo.EXPECT().
					Save(mock.Anything, tt.mocks.articleRepository.save_args_article).
					Return(tt.mocks.articleRepository.save_returns_error)
			}
			if !tt.configs.tag_bulkSave_should_be_skipped {
				tagRepo.EXPECT().
					BulkSave(mock.Anything, tt.mocks.tagRepository.bulkSave_args_tags).
					Return(tt.mocks.tagRepository.bulkSave_returns_error)
			}
			if tt.configs.errorHandler_handle_should_call {
				errorHandlerExpectations(t, errorHandler, errorHandlerExpectationsOption[CreateArticleResult]{
					HandleArgsError:      tt.mocks.errorHandler.handle_args_error,
					HandleArgsOptsLength: tt.mocks.errorHandler.handle_args_opts_length,
					HandleReturnsResult:  tt.mocks.errorHandler.handle_returns_result,
				})
			}

			a := NewCreateArticle(errorHandler, transaction, articleRepo, tagRepo, authServiec)
			got := a.Create(context.Background(), tt.args.param)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCreateArticleParam_ToDomain(t *testing.T) {
	type args struct {
		slug model.Slug
		user *authmodel.User
	}
	type wants struct {
		article *model.Article
		tags    []model.Tag
		err     error
	}
	type opts struct {
		nowFunc func() time.Time
	}

	var (
		now       = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		testData1 = struct {
			target CreateArticleParam
			args   args
			wants  wants
		}{
			target: CreateArticleParam{
				Title:       "title",
				Description: "description",
				Body:        "body",
				Tags: []string{
					"tag1", "tag2",
				},
			},
			args: args{
				slug: "slug",
				user: &authmodel.User{
					ID: "user1",
				},
			},
			wants: wants{
				article: &model.Article{
					Slug:   "slug",
					Author: "user1",
					Contents: model.ArticleContents{
						Title:       "title",
						Description: "description",
						Body:        "body",
					},
					Tags: []model.TagName{
						"tag1", "tag2",
					},
					CreatedAt: premitive.NewJSTTime(now),
					UpdatedAt: premitive.NewJSTTime(now),
				},
				tags: []model.Tag{
					{
						Tag:       "tag1",
						CreatedAt: premitive.NewJSTTime(now),
						UpdatedAt: premitive.NewJSTTime(now),
					},
					{
						Tag:       "tag2",
						CreatedAt: premitive.NewJSTTime(now),
						UpdatedAt: premitive.NewJSTTime(now),
					},
				},
			},
		}
	)

	tests := []struct {
		name   string
		target CreateArticleParam
		args   args
		wants  wants
		opts   opts
	}{
		{
			name:   "valid_params_should_return_expected_article",
			target: testData1.target,
			args:   testData1.args,
			wants:  testData1.wants,
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
		{
			name: "invalid_params_should_return_validation_error",
			target: CreateArticleParam{
				Title:       strings.Repeat("a", 10000),
				Description: testData1.target.Description,
				Body:        testData1.target.Body,
				Tags:        testData1.target.Tags,
			},
			args: testData1.args,
			wants: wants{
				err: derrors.Errors.Validation.Err,
			},
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
		{
			name: "invalid_tag_should_return_validation_error",
			target: CreateArticleParam{
				Title:       testData1.target.Title,
				Description: testData1.target.Description,
				Body:        testData1.target.Body,
				Tags: []string{
					strings.Repeat("a", 10000),
				},
			},
			args: testData1.args,
			wants: wants{
				err: derrors.Errors.Validation.Err,
			},
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
		{
			name:   "blank_slug_should_return_validation_error",
			target: testData1.target,
			args: args{
				slug: "",
				user: testData1.args.user,
			},
			wants: wants{
				err: derrors.Errors.Validation.Err,
			},
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
		{
			name:   "blank_user_should_return_validation_error",
			target: testData1.target,
			args: args{
				slug: testData1.args.slug,
				user: &authmodel.User{},
			},
			wants: wants{
				err: derrors.Errors.Validation.Err,
			},
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset := xtime.SetNowFunc(tt.opts.nowFunc)
			defer reset()

			artile, tags, err := tt.target.ToDomain(tt.args.slug, tt.args.user)
			assert.Equal(t, tt.wants.article, artile)
			assert.Equal(t, tt.wants.tags, tags)
			if !assert.ErrorIs(t, err, tt.wants.err) {
				t.Logf("%+v", err)
			}
		})
	}
}
