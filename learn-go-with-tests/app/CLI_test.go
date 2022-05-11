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

	assertPlayerWin(t, store, "Chris")
}
