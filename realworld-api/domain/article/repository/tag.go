package repository

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
)

type Tag interface {
	BulkSave(context.Context, []model.Tag) error
}
