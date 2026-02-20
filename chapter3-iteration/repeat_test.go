package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("Normal test, string length 1", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertRepeated(t, repeated, expected)
	})

	t.Run("Test with string of length greater than 1", func(t *testing.T) {
		repeated := Repeat("aaa", 3)
		expected := "aaaaaaaaa"

		assertRepeated(t, repeated, expected)
	})

	t.Run("Test with empty string", func(t *testing.T) {
		repeated := Repeat("", 10)
		expected := ""

		assertRepeated(t, repeated, expected)
	})

	t.Run("Test with negative integer", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""

		assertRepeated(t, repeated, expected)
	})

	t.Run("Test with 0 count", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertRepeated(t, repeated, expected)
	})

}

func assertRepeated(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("hello", 3)
	fmt.Println(repeated)
	// Output: hellohellohello
}
