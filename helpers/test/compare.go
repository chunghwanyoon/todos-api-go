package helpers

import "testing"

func Equals[T comparable](t *testing.T, want T, got T) {
	t.Helper()
	if want != got {
		t.Fatalf("TestFailedWithComparison: want %+v, got %+v", want, got)
	}
}
