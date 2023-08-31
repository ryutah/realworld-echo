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
    article_slug = any (sqlc.arg (slugs)::string[]);
