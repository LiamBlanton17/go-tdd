package main

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStringsEqual(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, gotErr := dictionary.Search("welcome")

		assertError(t, gotErr, ErrUnknownWord)
	})

}

func TestAdd(t *testing.T) {

	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		got, gotErr := dictionary.Search("test")
		want := "this is just a test"

		assertNoError(t, gotErr)
		assertStringsEqual(t, got, want)
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		gotAddErr := dictionary.Add(word, "this is a new test")
		got, gotSearchErr := dictionary.Search("test")

		assertError(t, gotAddErr, ErrWordAlreadyDefined)
		assertNoError(t, gotSearchErr)
		assertStringsEqual(t, got, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update existing", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		gotUpdateErr := dictionary.Update("test", "this is a new test")
		got, gotSearchErr := dictionary.Search("test")
		want := "this is a new test"

		assertNoError(t, gotUpdateErr)
		assertNoError(t, gotSearchErr)
		assertStringsEqual(t, got, want)
	})

	t.Run("update not existing", func(t *testing.T) {
		dictionary := Dictionary{}

		gotUpdateErr := dictionary.Update("welcome", "a formal greeting")
		_, gotSearchErr := dictionary.Search("welcome")

		assertError(t, gotUpdateErr, ErrUnknownWord)
		assertError(t, gotSearchErr, ErrUnknownWord)
	})

}

func TestDelete(t *testing.T) {

	t.Run("delete existing", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		gotDeleteErr := dictionary.Delete("test")
		_, gotSearchErr := dictionary.Search("test")

		assertNoError(t, gotDeleteErr)
		assertError(t, gotSearchErr, ErrUnknownWord)
	})

	t.Run("delete not existing", func(t *testing.T) {
		dictionary := Dictionary{}

		gotDeleteErr := dictionary.Delete("test")

		assertError(t, gotDeleteErr, ErrUnknownWord)
	})

}

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error")
	}

	assertStringsEqual(t, got.Error(), want.Error())
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("expected no error but got one")
	}
}
