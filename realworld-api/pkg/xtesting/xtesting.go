package xtesting

import (
	"testing"

	"github.com/cockroachdb/errors"
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
