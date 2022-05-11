package app_test

import (
	"strings"
	"testing"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCLI(store, in)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCLI(store, in)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Cleo")
	})
}
