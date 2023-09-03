-- name: ExistsFollow :one
select
    exists (
        select
            *
        from
            follow
        where
            user_id = $1
            and follwer_id = $2
    );


-- name: ExistsListFollow :many
with
    follow_exists as (
        select
            user_id,
            follwer_id,
            true as existance
        from
            follow
        where
            user_id = any (sqlc.arg (user_id)::text[])
            and follwer_id = $1
    )
select
    follow.user_id,
    follow.follwer_id,
    coalesce(follow_exists.existance, false) as existance
from
    follow
    left join follow_exists using (user_id, follwer_id)
where
    follow.user_id = any (sqlc.arg (user_id)::text[])
    and follow.follwer_id = $1;
