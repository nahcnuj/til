package main

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "This is just a test."}

	key := "test"
	got := Search(dict, key)
	want := "This is just a test."

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, key)
	}
}
