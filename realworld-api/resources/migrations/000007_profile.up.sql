create table
    profile (
        user_id user_id not null,
        bio long_text,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint profile_pkey primary key (user_id)
    );


comment on table "profile" is 'ユーザ情報';


comment on column profile.user_id is 'ユーザID';


comment on column profile.bio is '自己紹介';
