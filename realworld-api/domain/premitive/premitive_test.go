package premitive_test

import (
	"strings"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	. "github.com/ryutah/realworld-echo/realworld-api/domain/premitive"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtesting"
)

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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewSlug(test.slug)
			if diff := cmp.Diff(test.expected.slug, got.String()); diff != "" {
				xtesting.PrintDiff(t, "NewSlug", diff)
			}
			xtesting.CompareError(t, "NewSlug", test.expected.err, err)
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
			if diff := cmp.Diff(test.expected.title, got.String()); diff != "" {
				xtesting.PrintDiff(t, "NewTitle", diff)
			}
			xtesting.CompareError(t, "NewTitle", test.expected.err, err)
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
			if diff := cmp.Diff(test.expected.name, got.String()); diff != "" {
				xtesting.PrintDiff(t, "NewName", diff)
			}
			xtesting.CompareError(t, "NewName", test.expected.err, err)
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
			if diff := cmp.Diff(test.expected.email, got.String()); diff != "" {
				xtesting.PrintDiff(t, "NewEmail", diff)
			}

			if !xtesting.CompareError(t, "NewEmail", test.expected.err, err) {
				t.Log(errors.FlattenDetails(err))
			}
		})
	}
}
