package model

import (
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xvalidator"
)

type TagName premitive.ShortText

func NewTagName(name string) (TagName, error) {
	t, err := premitive.NewShortText(name)
	if err != nil {
		return "", err
	}
	tagName := TagName(t)
	if err := xvalidator.Validator().Var(tagName, "required"); err != nil {
		return "", errors.NewValidationError(0, err)
	}
	return tagName, nil
}

type Tag struct {
	Tag       TagName `validate:"required"`
	CreatedAt premitive.JSTTime
	UpdatedAt premitive.JSTTime
}

func NewTag(name TagName) (*Tag, error) {
	t := Tag{
		Tag:       name,
		CreatedAt: premitive.NewJSTTime(xtime.Now()),
		UpdatedAt: premitive.NewJSTTime(xtime.Now()),
	}

	if err := xvalidator.Validator().Struct(t); err != nil {
		return nil, errors.NewValidationError(0, err)
	}
	return &t, nil
}
