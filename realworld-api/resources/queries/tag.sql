-- name: CreateOrDoNothingTag :exec
insert into
    tag (name, created_at, updated_at)
values
    ($1, $2, $3)
on conflict (name) do nothing;
