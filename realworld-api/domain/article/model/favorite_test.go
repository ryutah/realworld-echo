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

func TestNewFavorite(t *testing.T) {
	type args struct {
		slug   Slug
		userID authmodel.UserID
	}
	type mocks struct {
		now func() time.Time
	}
	type wants struct {
		fav *Favorite
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
				fav: &Favorite{
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

			got, err := NewFavorite(tt.args.slug, tt.args.userID)
			assert.Equal(t, tt.want.fav, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}

func TestFavoriteSlice_IsFavorited(t *testing.T) {
	type args struct {
		userID authmodel.UserID
		slug   Slug
	}
	tests := []struct {
		name string
		f    FavoriteSlice
		args args
		want bool
	}{
		{
			name: "include_in_favarite_should_return_true",
			f: FavoriteSlice{
				{UserID: "user_id1", ArticleSlug: "slug1"},
				{UserID: "user_id1", ArticleSlug: "slug2"},
				{UserID: "user_id2", ArticleSlug: "slug1"},
			},
			args: args{
				userID: "user_id1",
				slug:   "slug2",
			},
			want: true,
		},
		{
			name: "not_include_in_favarite_should_return_false",
			f: FavoriteSlice{
				{UserID: "user_id1", ArticleSlug: "slug1"},
				{UserID: "user_id1", ArticleSlug: "slug2"},
				{UserID: "user_id2", ArticleSlug: "slug1"},
			},
			args: args{
				userID: "user_id2",
				slug:   "slug2",
			},
			want: false,
		},
		{
			name: "null_slice_should_return_false",
			f:    nil,
			args: args{
				userID: "user_id2",
				slug:   "slug2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f.IsFavorited(tt.args.userID, tt.args.slug)
			assert.Equal(t, tt.want, got)
		})
	}
}
