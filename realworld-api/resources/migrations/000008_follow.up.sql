create table
    follow (
        user_id user_id not null,
        follwer_id user_id not null,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint follow_pkey primary key (user_id, follwer_id)
    );


comment on table "follow" is 'フォロー';


comment on column follow.user_id is 'ユーザID';


comment on column follow.follwer_id is 'フォロワーユーザID';
