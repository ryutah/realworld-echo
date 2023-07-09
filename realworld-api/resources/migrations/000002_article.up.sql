CREATE TABLE article (
    slug uuid NOT NULL DEFAULT gen_random_uuid(),
    author user_id NOT NULL,
    body long_text NOT NULL,
    title short_text NOT NULL,
    description long_text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (slug)
);
COMMENT ON TABLE "article" IS '記事';
COMMENT ON COLUMN article.slug IS 'slug';
COMMENT ON COLUMN article.author IS '記事投稿者ID';
COMMENT ON COLUMN article.title IS '記事タイトル';
COMMENT ON COLUMN article.body IS '記事内容';
COMMENT ON COLUMN article.description IS '記事詳細情報';
