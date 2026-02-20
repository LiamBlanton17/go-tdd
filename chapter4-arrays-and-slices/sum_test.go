package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		assertSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {

	t.Run("Sum all first test", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertSumAll(t, got, want)
	})

}

func TestSumAllTails(t *testing.T) {

	t.Run("Sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertSumallTails(t, got, want)
	})

	t.Run("Sum with an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		assertSumallTails(t, got, want)
	})

}

func assertSum(t testing.TB, got, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func assertSumAll(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertSumallTails(t testing.TB, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
