package model

import (
	"time"

	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
)

type CommentID premitive.UID

type Comment struct {
	ID        CommentID
	ArticleID ArticleID
	Author    UserID
	Body      premitive.LongText
	CreatedAt time.Time
	UpdatedAt time.Time
}
