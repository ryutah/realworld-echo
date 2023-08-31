create table
    article_tag (
        article_slug uuid not null,
        tag_name short_text not null,
        created_at timestamp with time zone not null default current_timestamp,
        updated_at timestamp with time zone not null default current_timestamp,
        constraint article_tag_pkey primary key (article_slug, tag_name),
        constraint fk_article_tag_article foreign key (article_slug) references article (slug),
        constraint fk_article_tag_tag foreign key (tag_name) references tag (name)
    );


comment on table "article_tag" is 'タグ情報';


comment on column article_tag.article_slug is '記事slug';


comment on column article_tag.tag_name is 'タグ名';
