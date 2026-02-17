package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Testing hello world with a name", func(t *testing.T) {
		got := Hello("Liam", "")
		want := "Hello, Liam!"

		assertHelloTest(t, got, want)
	})

	t.Run("Testing hello world with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertHelloTest(t, got, want)
	})

	t.Run("Testing hello world with Spanish", func(t *testing.T) {
		got := Hello("Liam", "Spanish")
		want := "Hola, Liam!"

		assertHelloTest(t, got, want)
	})

	t.Run("Testing hello world with French", func(t *testing.T) {
		got := Hello("Liam", "French")
		want := "Bonjour, Liam!"

		assertHelloTest(t, got, want)
	})
}

func assertHelloTest(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
