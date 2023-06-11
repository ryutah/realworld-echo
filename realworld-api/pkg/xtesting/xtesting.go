package xtesting

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/golang/mock/gomock"
)

func PrintDiff(t *testing.T, name string, diff string) {
	t.Helper()
	t.Errorf("%v: (-want, got)\n%s", name, diff)
}

func CompareError(t *testing.T, name string, want, got error) (isMatch bool) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("Error(%s) = %v, want %v", name, got, want)
		return false
	}
	return true
}

type errorMatcher struct {
	want error
}

func MatchError(want error) gomock.Matcher {
	return &errorMatcher{
		want: want,
	}
}

// Matches returns whether x is a match.
func (e *errorMatcher) Matches(x any) bool {
	got, ok := x.(error)
	if !ok {
		return false
	}
	return errors.Is(got, e.want)
}

// String describes what the matcher matches.
func (e *errorMatcher) String() string {
	return fmt.Sprintf("%#v", e.want)
}
