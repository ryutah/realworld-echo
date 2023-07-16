package model_test

import (
	"testing"
	"time"

	. "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/stretchr/testify/assert"
)

func TestNewArticleFavorite(t *testing.T) {
	type args struct {
		slug   Slug
		userID authmodel.UserID
	}
	type mocks struct {
		now func() time.Time
	}
	type wants struct {
		fav *ArticleFavorite
		err error
	}

	var (
		now     = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		nowFunc = func() time.Time { return now }
	)

	tests := []struct {
		name  string
		args  args
		mocks mocks
		want  wants
	}{
		{
			name: "valid_slug_and_user_id_should_return_expected_favorite",
			args: args{
				slug:   "slug",
				userID: "user_id",
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				fav: &ArticleFavorite{
					ArticleSlug: "slug",
					UserID:      "user_id",
					CreatedAt:   premitive.NewJSTTime(now),
					UpdateAt:    premitive.NewJSTTime(now),
				},
				err: nil,
			},
		},
		{
			name: "blank_slug_should_return_validation_error",
			args: args{
				slug:   "",
				userID: "user_id",
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
		{
			name: "blank_user_id_should_return_validation_error",
			args: args{
				slug:   "slug",
				userID: "",
			},
			mocks: mocks{
				now: nowFunc,
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset := xtime.SetNowFunc(tt.mocks.now)
			defer reset()

			got, err := NewArticleFavorite(tt.args.slug, tt.args.userID)
			assert.Equal(t, tt.want.fav, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
