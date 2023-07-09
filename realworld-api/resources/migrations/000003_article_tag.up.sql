CREATE TABLE article_tag (
    article_slug uuid NOT NULL,
    tag short_text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (article_slug, tag),
    CONSTRAINT fk_article_tag_article FOREIGN KEY (
        article_slug
    ) REFERENCES article (slug)
);
COMMENT ON TABLE "article_tag" IS 'タグ情報';
COMMENT ON COLUMN article_tag.article_slug IS '記事slug';
COMMENT ON COLUMN article_tag.tag IS 'タグ';
