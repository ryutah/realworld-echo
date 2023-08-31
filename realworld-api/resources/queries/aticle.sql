-- name: GetArticle :one
select
    *
from
    article
where
    slug = $1
limit
    1;


-- name: UpsertArticle :exec
insert into
    article (
        slug,
        author,
        body,
        title,
        description,
        created_at,
        updated_at
    )
values
    ($1, $2, $3, $4, $5, $6, $7)
on conflict (slug) do
update
set
    author = $2,
    body = $3,
    title = $4,
    description = $5,
    updated_at = $7;
