package app

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	store := &StubPlayerStore{}
	cli := &CLI{store, in}
	cli.PlayPoker()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't call")
	}

	got := store.winCalls[0]
	want := "Chris"

	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
