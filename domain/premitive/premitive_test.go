package premitive_test

import (
	"strings"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	derrors "github.com/ryutah/realworld-echo/domain/errors"
	. "github.com/ryutah/realworld-echo/domain/premitive"
)

func TestSlug(t *testing.T) {
	type expected struct {
		slug Slug
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
				slug: Slug("valid"),
				err:  nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewSlug(test.slug)
			if diff := cmp.Diff(test.expected.slug, got); diff != "" {
				t.Errorf("NewSlug(%q) got diff: %s", test.slug, diff)
			}
			if !errors.Is(err, test.expected.err) {
				t.Errorf("NewSlug(%q) got expected error: %v, got %v", test.slug, test.expected.err, err)
				t.Log(errors.FlattenDetails(err))
			}
		})
	}
}

func TestTitle(t *testing.T) {
	type expected struct {
		title Title
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
				title: Title("valid"),
				err:   nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewTitle(test.title)
			if diff := cmp.Diff(test.expected.title, got); diff != "" {
				t.Errorf("NewTitle(%q) got diff: %s", test.title, diff)
			}
			if !errors.Is(err, test.expected.err) {
				t.Errorf("NewTitle(%q) got expected error: %v, got %v", test.title, test.expected.err, err)
				t.Log(errors.FlattenDetails(err))
			}
		})
	}
}

func TestName(t *testing.T) {
	type expected struct {
		name Name
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
				name: Name("valid"),
				err:  nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewName(test.n)
			if diff := cmp.Diff(test.expected.name, got); diff != "" {
				t.Errorf("NewName(%q) got diff: %s", test.n, diff)
			}
			if !errors.Is(err, test.expected.err) {
				t.Errorf("NewName(%q) got expected error: %v, got %v", test.n, test.expected.err, err)
				t.Log(errors.FlattenDetails(err))
			}
		})
	}
}

func TestEmail(t *testing.T) {
	type expected struct {
		email Email
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
				email: Email("example@gmail.com"),
				err:   nil,
			},
		},
		{
			name:  "valid_max_length_email",
			email: strings.Repeat("a", 244) + "@gmail.com",
			expected: expected{
				email: Email(strings.Repeat("a", 244) + "@gmail.com"),
				err:   nil,
			},
		},
		{
			name:  "invalid_email",
			email: "invalid_email_address",
			expected: expected{
				email: Email(""),
				err:   derrors.Errors.ErrValidation.Err,
			},
		},
		{
			name:  "invalid_length_email",
			email: strings.Repeat("a", 245) + "@gmail.com",
			expected: expected{
				email: Email(""),
				err:   derrors.Errors.ErrValidation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewEmail(test.email)
			if diff := cmp.Diff(test.expected.email, got); diff != "" {
				t.Errorf("NewEmail(%q) got diff: %s", test.email, diff)
			}
			if !errors.Is(err, test.expected.err) {
				t.Errorf("NewEmail(%q) got expected error: %v, got %v", test.email, test.expected.err, err)
				t.Log(errors.FlattenDetails(err))
			}
		})
	}
}
