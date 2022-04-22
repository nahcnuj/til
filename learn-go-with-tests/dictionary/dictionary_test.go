package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "This is just a test."}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "This is just a test."

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		value := "this is just a test"

		dict.Add(key, value)
		assertDefinition(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}

		err := dict.Add(key, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dict := Dictionary{key: value}
	newValue := "new definition"

	dict.Update(key, newValue)
	assertDefinition(t, dict, key, newValue)
}

func assertDefinition(t testing.TB, dict Dictionary, key, want string) {
	t.Helper()

	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
