package model_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	type args struct {
		tag TagName
	}
	type wants struct {
		tag *Tag
		err error
	}
	type opts struct {
		nowFunc func() time.Time
	}

	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want wants
		opts opts
	}{
		{
			name: "valid_tag_should_return_expecte_tag",
			args: args{
				tag: "valid_tag",
			},
			want: wants{
				tag: &Tag{
					Tag:       "valid_tag",
					CreatedAt: premitive.NewJSTTime(now),
					UpdatedAt: premitive.NewJSTTime(now),
				},
				err: nil,
			},
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
		{
			name: "invali_tag_should_return_validation_error",
			args: args{
				tag: "",
			},
			want: wants{
				tag: nil,
				err: errors.Errors.Validation.Err,
			},
			opts: opts{
				nowFunc: func() time.Time { return now },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset := xtime.SetNowFunc(tt.opts.nowFunc)
			defer reset()

			got, err := NewTag(tt.args.tag)

			assert.Equal(t, tt.want.tag, got)
			if !assert.ErrorIs(t, err, tt.want.err) {
				t.Logf("%+v", err)
			}
		})
	}
}

func TestNewTagName(t *testing.T) {
	type args struct {
		name string
	}
	type wants struct {
		name TagName
		err  error
	}

	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "valid_tag_should_return_expected_tag",
			args: args{
				name: "valid_tag",
			},
			want: wants{
				name: "valid_tag",
			},
		},
		{
			name: "invalid_tag_should_return_validation_error",
			args: args{
				name: strings.Repeat("a", 10000),
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
		{
			name: "blank_tag_should_return_validation_error",
			args: args{
				name: "",
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTagName(tt.args.name)
			assert.Equal(t, tt.want.name, got)
			if !assert.ErrorIs(t, err, tt.want.err) {
				t.Logf("%+v", err)
			}
		})
	}
}
