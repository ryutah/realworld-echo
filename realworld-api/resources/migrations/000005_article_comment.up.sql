CREATE TABLE article_comment (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    article_slug uuid NOT NULL,
    author user_id NOT NULL,
    body long_text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (id),
    CONSTRAINT fk_article_comment_article FOREIGN KEY (
        article_slug
    ) REFERENCES article (slug)
);
COMMENT ON TABLE "article_comment" IS '記事コメント';
COMMENT ON COLUMN article_comment.id IS 'コメントID';
COMMENT ON COLUMN article_comment.article_slug IS '記事slug';
COMMENT ON COLUMN article_comment.author IS '投稿者';
COMMENT ON COLUMN article_comment.body IS 'コメント内容';
