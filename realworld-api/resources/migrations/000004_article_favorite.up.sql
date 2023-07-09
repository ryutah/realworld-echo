CREATE TABLE article_favorite (
    article_slug uuid NOT NULL,
    user_id user_id NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (article_slug, user_id),
    CONSTRAINT fk_article_favorite_article FOREIGN KEY (
        article_slug
    ) REFERENCES article (slug)
);
COMMENT ON TABLE "article_favorite" IS '記事フォロー';
COMMENT ON COLUMN article_favorite.article_slug IS '記事slug';
COMMENT ON COLUMN article_favorite.user_id IS 'ユーザID';
