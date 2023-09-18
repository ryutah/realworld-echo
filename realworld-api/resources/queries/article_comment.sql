-- name: GetArticleComment :one
select
    *
from
    article_comment
where
    id = $1;


-- name: ListArticleCommentBySlug :one
select
    *
from
    article_comment
where
    article_slug = $1;


-- name: UpsertArticleComment :exec
insert into
    article_comment (
        id,
        article_slug,
        author,
        body,
        created_at,
        updated_at
    )
values
    ($1, $2, $3, $4, $5, $6)
on conflict (id) do
update
set
    body = $4,
    updated_at = $6;


-- name: DeleteArticleComment :exec
delete from article_comment
where
    id = $1;
