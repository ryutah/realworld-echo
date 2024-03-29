// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: article_favorite.sql

package gen

import (
	"context"

	"github.com/google/uuid"
)

const CountArticleFavorite = `-- name: CountArticleFavorite :one
select
    count(*) as count
from
    article_favorite
where
    article_slug = $1
`

func (q *Queries) CountArticleFavorite(ctx context.Context, articleSlug uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, CountArticleFavorite, articleSlug)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CountListArticleFavorite = `-- name: CountListArticleFavorite :many
select
    article_slug,
    count(*) as count
from
    article_favorite
where
    article_slug = any ($1::uuid[])
group by
    article_slug
`

type CountListArticleFavoriteRow struct {
	ArticleSlug uuid.UUID `db:"article_slug"`
	Count       int64     `db:"count"`
}

func (q *Queries) CountListArticleFavorite(ctx context.Context, slugs []uuid.UUID) ([]CountListArticleFavoriteRow, error) {
	rows, err := q.db.Query(ctx, CountListArticleFavorite, slugs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CountListArticleFavoriteRow
	for rows.Next() {
		var i CountListArticleFavoriteRow
		if err := rows.Scan(&i.ArticleSlug, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ExistsArticleFavorite = `-- name: ExistsArticleFavorite :one
select
    exists (
        select
            article_slug, user_id, created_at, updated_at
        from
            article_favorite
        where
            article_slug = $1
            and user_id = $2
    )
`

type ExistsArticleFavoriteParams struct {
	ArticleSlug uuid.UUID `db:"article_slug"`
	UserID      string    `db:"user_id"`
}

func (q *Queries) ExistsArticleFavorite(ctx context.Context, arg ExistsArticleFavoriteParams) (bool, error) {
	row := q.db.QueryRow(ctx, ExistsArticleFavorite, arg.ArticleSlug, arg.UserID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const ExistsListArtileFavorite = `-- name: ExistsListArtileFavorite :many
with
    favorites_exists as (
        select
            article_slug,
            user_id,
            true as existance
        from
            article_favorite
        where
            article_slug = any ($2::uuid[])
            and user_id = $1
    )
select
    article_favorite.article_slug,
    article_favorite.user_id,
    coalesce(favorites_exists.existance, false) as existance
from
    article_favorite
    left join favorites_exists using (article_slug, user_id)
where
    article_favorite.article_slug = any ($2::uuid[])
    and article_favorite.user_id = $1
`

type ExistsListArtileFavoriteParams struct {
	UserID string      `db:"user_id"`
	Slugs  []uuid.UUID `db:"slugs"`
}

type ExistsListArtileFavoriteRow struct {
	ArticleSlug uuid.UUID `db:"article_slug"`
	UserID      string    `db:"user_id"`
	Existance   bool      `db:"existance"`
}

func (q *Queries) ExistsListArtileFavorite(ctx context.Context, arg ExistsListArtileFavoriteParams) ([]ExistsListArtileFavoriteRow, error) {
	rows, err := q.db.Query(ctx, ExistsListArtileFavorite, arg.UserID, arg.Slugs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExistsListArtileFavoriteRow
	for rows.Next() {
		var i ExistsListArtileFavoriteRow
		if err := rows.Scan(&i.ArticleSlug, &i.UserID, &i.Existance); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
