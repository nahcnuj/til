package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "This is just a test."}

	t.Run("known word", func(t *testing.T) {
		got := dict.Search("test")
		want := "This is just a test."

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		want := "could not find the word you were looking for"
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
