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
		{
			name:  "max_length_title",
			title: strings.Repeat("a", 255),
			expected: expected{
				title: strings.Repeat("a", 255),
				err:   nil,
			},
		},
		{
			name:  "invalid_length_title",
			title: strings.Repeat("a", 256),
			expected: expected{
				err: derrors.Errors.Validation.Err,
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

func TestLongText(t *testing.T) {
	type expected struct {
		text string
		err  error
	}

	tests := []struct {
		name     string
		text     string
		expected expected
	}{
		{
			name: "valid_longtext",
			text: "long_text",
			expected: expected{
				text: "long_text",
				err:  nil,
			},
		},
		{
			name: "valid_max_length_longtext",
			text: strings.Repeat("a", 5000),
			expected: expected{
				text: strings.Repeat("a", 5000),
				err:  nil,
			},
		},
		{
			name: "invalid_length_longtext",
			text: strings.Repeat("a", 5001),
			expected: expected{
				text: "",
				err:  derrors.Errors.Validation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewLongText(test.text)
			assert.Equal(t, test.expected.text, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}

func TestShortText(t *testing.T) {
	type expected struct {
		text string
		err  error
	}

	tests := []struct {
		name     string
		text     string
		expected expected
	}{
		{
			name: "valid_shorttext",
			text: "short_text",
			expected: expected{
				text: "short_text",
				err:  nil,
			},
		},
		{
			name: "valid_max_length_shorttext",
			text: strings.Repeat("a", 255),
			expected: expected{
				text: strings.Repeat("a", 255),
				err:  nil,
			},
		},
		{
			name: "invalid_length_shorttext",
			text: strings.Repeat("a", 256),
			expected: expected{
				text: "",
				err:  derrors.Errors.Validation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewShortText(test.text)
			assert.Equal(t, test.expected.text, got.String())
			assert.ErrorIs(t, err, test.expected.err)
		})
	}
}

func TestURL(t *testing.T) {
	type expected struct {
		url string
		err error
	}

	tests := []struct {
		name     string
		url      string
		expected expected
	}{
		{
			name: "valid_url",
			url:  "https://test.com",
			expected: expected{
				url: "https://test.com",
				err: nil,
			},
		},
		{
			name: "invalid_url",
			url:  "bad url",
			expected: expected{
				url: "",
				err: derrors.Errors.Validation.Err,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewURL(test.url)
			assert.Equal(t, test.expected.url, got.String())
			if !assert.ErrorIs(t, err, test.expected.err) {
				t.Logf("error: %+v", err)
			}
		})
	}
}

func TestNewCount(t *testing.T) {
	type args struct {
		i uint
	}
	type wants struct {
		uint uint
		int  int
	}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "valid_count",
			args: args{
				i: 100,
			},
			wants: wants{
				uint: 100,
				int:  100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCount(tt.args.i)
			assert.Equal(t, tt.wants.uint, got.Uint())
			assert.Equal(t, tt.wants.int, got.Int())
		})
	}
}
