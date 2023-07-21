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

func TestNewAccount(t *testing.T) {
	type args struct {
		email string
	}
	type wants struct {
		account *Account
		err     error
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "valid_email_should_return_expected_account",
			args: args{
				email: "example@gmail.com",
			},
			wants: wants{
				account: &Account{
					Email: "example@gmail.com",
				},
			},
		},
		{
			name: "invalid_email_should_return_validation_error",
			args: args{
				email: "invalid",
			},
			wants: wants{
				err: errors.Errors.Validation.Err,
			},
		},
		{
			name: "blank_email_should_return_validation_error",
			args: args{
				email: "",
			},
			wants: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccount(tt.args.email)
			assert.Equal(t, tt.wants.account, got)
			assert.ErrorIs(t, err, tt.wants.err)
		})
	}
}

func TestNewProfile(t *testing.T) {
	type args struct {
		username string
		image    string
	}
	type wants struct {
		profile *Profile
		err     error
	}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "valid_username_and_image_should_return_expected_profile",
			args: args{
				username: "username",
				image:    "http://example.com/image.png",
			},
			wants: wants{
				profile: &Profile{
					Username: "username",
					Image:    "http://example.com/image.png",
				},
			},
		},
		{
			name: "invalid_username_should_return_validation_error",
			args: args{
				username: strings.Repeat("a", 10000),
				image:    "http://example.com/image.png",
			},
			wants: wants{
				err: errors.Errors.Validation.Err,
			},
		},
		{
			name: "invalid_image_should_return_validation_error",
			args: args{
				image: "not_url",
			},
			wants: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProfile(tt.args.username, tt.args.image)
			assert.Equal(t, tt.wants.profile, got)
			assert.ErrorIs(t, err, tt.wants.err)
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		id      UserID
		account Account
		profile Profile
	}
	type wants struct {
		user *User
		err  error
	}
	var (
		account = Account{
			Email: "example@gmail.com",
		}
		profile = Profile{
			Username: "username",
			Image:    "http://example.com/image.png",
		}
	)

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "valid_arguments_should_return_expected_user",
			args: args{
				id:      "uid",
				account: account,
				profile: profile,
			},
			wants: wants{
				user: &User{
					ID:      "uid",
					Account: account,
					Profile: profile,
				},
			},
		},
		{
			name: "blank_id_should_return_validation_error",
			args: args{
				id:      "",
				account: account,
				profile: profile,
			},
			wants: wants{
				err: errors.Errors.Validation.Err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.id, tt.args.account, tt.args.profile)
			assert.Equal(t, tt.wants.user, got)
			assert.ErrorIs(t, err, tt.wants.err)
		})
	}
}
