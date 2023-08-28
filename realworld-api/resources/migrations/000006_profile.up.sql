CREATE TABLE profile (
    user_id uuid NOT NULL,
    bio long_text,
    PRIMARY KEY (user_id),
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp
);
COMMENT ON TABLE "profile" IS 'ユーザ情報';
COMMENT ON COLUMN profile.user_id IS 'ユーザID';
COMMENT ON COLUMN profile.bio IS '自己紹介';
