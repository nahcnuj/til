package app_test

import (
	"bytes"
	"io"
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
	t.Run("start a game with 3 players and finish it with Chris as the winner", func(t *testing.T) {
		in := userSends("3", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, app.PlayerPrompt)
		assertNumberOfPlayers(t, game.StartedWith, 3)
		assertWinner(t, game.FinishedWith, "Chris")
	})

	t.Run("start a game with 8 players and record Cleo as the winner", func(t *testing.T) {
		in := userSends("8", "Cleo wins")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, app.PlayerPrompt)
		assertNumberOfPlayers(t, game.StartedWith, 8)
		assertWinner(t, game.FinishedWith, "Cleo")
	})

	t.Run("print an error if a non numeric value is entered", func(t *testing.T) {
		in := userSends("Pies")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.CalledStart {
			t.Error("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, app.PlayerPrompt, app.BadPlayerInputError)
	})
}

func userSends(inputs ...string) io.Reader {
	s := ""
	for _, input := range inputs {
		s += input + "\n"
	}
	return strings.NewReader(s)
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

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	got := stdout.String()
	want := strings.Join(messages, "")
	if got != want {
		t.Errorf("got %q sent to user, but want %+v", got, messages)
	}
}
