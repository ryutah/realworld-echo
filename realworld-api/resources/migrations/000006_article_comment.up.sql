create table
    article_comment (
        id uuid not null default gen_random_uuid (),
        article_slug uuid not null,
        author user_id not null,
        body long_text not null,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint article_comment_pkey primary key (id),
        constraint fk_article_comment_article foreign key (article_slug) references article (slug)
    );


comment on table "article_comment" is '記事コメント';


comment on column article_comment.id is 'コメントID';


comment on column article_comment.article_slug is '記事slug';


comment on column article_comment.author is '投稿者';


comment on column article_comment.body is 'コメント内容';
