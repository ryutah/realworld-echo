// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package gen

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	BulkCreateTagOrDoNothing(ctx context.Context, arg BulkCreateTagOrDoNothingParams) error
	CountArticleFavorite(ctx context.Context, articleSlug uuid.UUID) (int64, error)
	CountListArticleFavorite(ctx context.Context, slugs []uuid.UUID) ([]CountListArticleFavoriteRow, error)
	CreateArticleTag(ctx context.Context, arg []CreateArticleTagParams) (int64, error)
	CreateOrDoNothingTag(ctx context.Context, arg CreateOrDoNothingTagParams) error
	DeleteArticleTagBySlug(ctx context.Context, articleSlug uuid.UUID) error
	ExistsArticleFavorite(ctx context.Context, arg ExistsArticleFavoriteParams) (bool, error)
	ExistsFollow(ctx context.Context, arg ExistsFollowParams) (bool, error)
	ExistsListArtileFavorite(ctx context.Context, arg ExistsListArtileFavoriteParams) ([]ExistsListArtileFavoriteRow, error)
	ExistsListFollow(ctx context.Context, arg ExistsListFollowParams) ([]ExistsListFollowRow, error)
	GetArticle(ctx context.Context, slug uuid.UUID) (Article, error)
	ListArticleTags(ctx context.Context, slugs []uuid.UUID) ([]ListArticleTagsRow, error)
	UpsertArticle(ctx context.Context, arg UpsertArticleParams) error
}

var _ Querier = (*Queries)(nil)
