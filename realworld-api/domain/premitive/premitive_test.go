package premitive_test

import (
	"strings"
	"testing"

	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	. "github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/stretchr/testify/assert"
)

func TestUID(t *testing.T) {
	type expected struct {
		uid string
		err error
	}

	tests := []struct {
		name     string
		uid      string
		expected expected
	}{
		{
			name: "valid_uid",
			uid:  "valid",
			expected: expected{
				uid: "valid",
				err: nil,
			},
		},
		{
			name: "invalid_length_uid",
			uid:  strings.Repeat("a", 256),
			expected: expected{
				uid: "",
				err: derrors.Errors.Validation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewUID(test.uid)
			assert.Equal(t, test.expected.uid, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}

func TestSlug(t *testing.T) {
	type expected struct {
		slug string
		err  error
	}

	tests := []struct {
		name     string
		slug     string
		expected expected
	}{
		{
			name: "valid_slug",
			slug: "valid",
			expected: expected{
				slug: "valid",
				err:  nil,
			},
		},
		{
			name: "invalid_length_slug",
			slug: strings.Repeat("a", 51),
			expected: expected{
				slug: "",
				err:  derrors.Errors.Validation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewSlug(test.slug)
			assert.Equal(t, test.expected.slug, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}

func TestTitle(t *testing.T) {
	type expected struct {
		title string
		err   error
	}

	tests := []struct {
		name     string
		title    string
		expected expected
	}{
		{
			name:  "valid_title",
			title: "valid",
			expected: expected{
				title: "valid",
				err:   nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewTitle(test.title)
			assert.Equal(t, test.expected.title, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}

func TestName(t *testing.T) {
	type expected struct {
		name string
		err  error
	}

	tests := []struct {
		name     string
		n        string
		expected expected
	}{
		{
			name: "valid_title",
			n:    "valid",
			expected: expected{
				name: "valid",
				err:  nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewName(test.n)
			assert.Equal(t, test.expected.name, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}

func TestEmail(t *testing.T) {
	type expected struct {
		email string
		err   error
	}

	tests := []struct {
		name     string
		email    string
		expected expected
	}{
		{
			name:  "valid_email",
			email: "example@gmail.com",
			expected: expected{
				email: "example@gmail.com",
				err:   nil,
			},
		},
		{
			name:  "valid_max_length_email",
			email: strings.Repeat("a", 244) + "@gmail.com",
			expected: expected{
				email: strings.Repeat("a", 244) + "@gmail.com",
				err:   nil,
			},
		},
		{
			name:  "invalid_email",
			email: "invalid_email_address",
			expected: expected{
				email: "",
				err:   derrors.Errors.Validation.Err,
			},
		},
		{
			name:  "invalid_length_email",
			email: strings.Repeat("a", 245) + "@gmail.com",
			expected: expected{
				email: "",
				err:   derrors.Errors.Validation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewEmail(test.email)
			assert.Equal(t, test.expected.email, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}
