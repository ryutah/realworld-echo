package model

import (
	"time"

	"github.com/ryutah/realworld-echo/domain/premitive"
)

type Article struct {
	Slug        premitive.Slug
	Title       premitive.Title
	Description premitive.Description
	Body        premitive.LongBody
	Author      *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
