-- name: CreateOrDoNothingTag :exec
insert into
    tag (name, created_at, updated_at)
values
    ($1, $2, $3)
on conflict (name) do nothing;


-- name: BulkCreateTagOrDoNothing :exec
insert into
    tag (name, created_at, updated_at)
select
    unnest(sqlc.arg (name)::text[]) as name,
    unnest(sqlc.arg (created_at)::timestamp[]) as created_at,
    unnest(sqlc.arg (updated_at)::timestamp[]) as updated_at
on conflict (name) do nothing;
