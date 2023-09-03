-- name: CountArticleFavorite :one
select
    count(*) as count
from
    article_favorite
where
    article_slug = $1;


-- name: CountListArticleFavorite :many
select
    article_slug,
    count(*) as count
from
    article_favorite
where
    article_slug = any (sqlc.arg (slugs)::uuid[])
group by
    article_slug;


-- name: ExistsArticleFavorite :one
select
    exists (
        select
            *
        from
            article_favorite
        where
            article_slug = $1
            and user_id = $2
    );


-- name: ExistsListArtileFavorite :many
with
    favorites_exists as (
        select
            article_slug,
            user_id,
            true as existance
        from
            article_favorite
        where
            article_slug = any (sqlc.arg (slugs)::uuid[])
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
    article_favorite.article_slug = any (sqlc.arg (slugs)::uuid[])
    and article_favorite.user_id = $1;
