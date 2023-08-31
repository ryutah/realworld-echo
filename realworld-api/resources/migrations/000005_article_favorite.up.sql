create table
    article_favorite (
        article_slug uuid not null,
        user_id user_id not null,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint article_favorite_pkey primary key (article_slug, user_id),
        constraint fk_article_favorite_article foreign key (article_slug) references article (slug)
    );


comment on table "article_favorite" is '記事フォロー';


comment on column article_favorite.article_slug is '記事slug';


comment on column article_favorite.user_id is 'ユーザID';
