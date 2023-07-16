package model_test

import (
	"strings"
	"testing"

	. "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"
	"github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewUserID(t *testing.T) {
	type args struct {
		s string
	}
	type wants struct {
		id  string
		err error
	}

	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "valid_id_should_return_expected_id",
			args: args{
				s: "uid",
			},
			want: wants{
				id: "uid",
			},
		},
		{
			name: "invalid_id_should_return_validation_error",
			args: args{
				s: strings.Repeat("a", 10000),
			},
			want: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserID(tt.args.s)

			assert.Equal(t, tt.want.id, got.String())
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
