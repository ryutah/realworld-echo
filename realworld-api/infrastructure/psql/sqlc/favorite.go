package sqlc

import (
	"github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
)

type Favorite struct {
	repository.Favorite
}

var _ repository.Favorite = (*Favorite)(nil)

func NewFavorite() *Favorite {
	return &Favorite{}
}
