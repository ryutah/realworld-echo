-- name: ListArticleTags :many
select
    article_tag.article_slug,
    tag.name,
    tag.created_at,
    tag.updated_at
from
    article_tag
    inner join tag on tag.name = article_tag.tag_name
where
    article_slug = any (sqlc.arg (slugs)::uuid[]);


-- name: DeleteArticleTagBySlug :exec
delete from article_tag
where
    article_slug = $1;


-- name: CreateArticleTag :copyfrom
insert into
    article_tag (article_slug, tag_name, created_at, updated_at)
values
    ($1, $2, $3, $4);
