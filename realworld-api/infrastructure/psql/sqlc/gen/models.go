// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package gen

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// 記事
type Article struct {
	// slug
	Slug uuid.UUID `db:"slug"`
	// 記事投稿者ID
	Author string `db:"author"`
	// 記事内容
	Body string `db:"body"`
	// 記事タイトル
	Title string `db:"title"`
	// 記事詳細情報
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// 記事コメント
type ArticleComment struct {
	// コメントID
	ID uuid.UUID `db:"id"`
	// 記事slug
	ArticleSlug uuid.UUID `db:"article_slug"`
	// 投稿者
	Author string `db:"author"`
	// コメント内容
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 記事フォロー
type ArticleFavorite struct {
	// 記事slug
	ArticleSlug uuid.UUID `db:"article_slug"`
	// ユーザID
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// タグ情報
type ArticleTag struct {
	// 記事slug
	ArticleSlug uuid.UUID `db:"article_slug"`
	// タグ名
	TagName   string    `db:"tag_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ユーザ情報
type Profile struct {
	// ユーザID
	UserID uuid.UUID `db:"user_id"`
	// 自己紹介
	Bio       sql.NullString `db:"bio"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

// タグ
type Tag struct {
	// タグ名
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}