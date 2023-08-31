create table
    article (
        slug uuid not null default gen_random_uuid (),
        author user_id not null,
        body long_text not null,
        title short_text not null,
        description long_text not null,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint article_pkey primary key (slug)
    );


comment on table "article" is '記事';


comment on column article.slug is 'slug';


comment on column article.author is '記事投稿者ID';


comment on column article.title is '記事タイトル';


comment on column article.body is '記事内容';


comment on column article.description is '記事詳細情報';
