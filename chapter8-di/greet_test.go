package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		assertStringsEqual(t, got, want)
	})

}

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
