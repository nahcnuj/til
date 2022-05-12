package app_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

var dummyStdOut = &bytes.Buffer{}

type SpyGame struct {
	CalledStart  bool
	StartedWith  int
	FinishedWith string
}

func (g *SpyGame) Start(numberOfPlayers int) {
	g.CalledStart = true
	g.StartedWith = numberOfPlayers
}

func (g *SpyGame) Finish(winner string) {
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		game := &SpyGame{}

		cli := app.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertWinner(t, game.FinishedWith, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &SpyGame{}

		cli := app.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertWinner(t, game.FinishedWith, "Cleo")
	})

	t.Run("start with 5 players and finish", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		game := &SpyGame{}

		cli := app.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertNumberOfPlayers(t, game.StartedWith, 5)
		assertWinner(t, game.FinishedWith, "Chris")
	})

	t.Run("prompt the user to enter the number of players first", func(t *testing.T) {
		in := strings.NewReader("7\n")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := app.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("wrong prompt, got %q, want %q", gotPrompt, wantPrompt)
		}

		assertNumberOfPlayers(t, game.StartedWith, 7)
	})

	t.Run("print an error if a non numeric value is entered", func(t *testing.T) {
		in := strings.NewReader("Pies\n")

		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.CalledStart {
			t.Error("game should not have started")
		}
	})
}

func assertNumberOfPlayers(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d players, but got %d", want, got)
	}
}

func assertWinner(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected winner %s, but got %q", want, got)
	}
}
