package app_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("start a game with 3 players and finish it with Chris as the winner", func(t *testing.T) {
		in := userSends("3", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &app.SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, app.PlayerPrompt)
		app.AssertGameStartedWith(t, game, 3)
		app.AssertGameFinishedWith(t, game, "Chris")
	})

	t.Run("start a game with 8 players and record Cleo as the winner", func(t *testing.T) {
		in := userSends("8", "Cleo wins")
		stdout := &bytes.Buffer{}
		game := &app.SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, app.PlayerPrompt)
		app.AssertGameStartedWith(t, game, 8)
		app.AssertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("print an error if a non numeric value is entered", func(t *testing.T) {
		in := userSends("Pies")
		stdout := &bytes.Buffer{}
		game := &app.SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		app.AssertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, app.PlayerPrompt, app.BadPlayerInputError)
	})

	t.Run("print an error if winner could not be parsed", func(t *testing.T) {
		in := userSends("5", "Lloyd is a killer")
		stdout := &bytes.Buffer{}
		game := &app.SpyGame{}

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.CalledFinish {
			t.Error("game should not have finished")
		}
		assertMessagesSentToUser(t, stdout, app.PlayerPrompt, app.BadWinnerInputError)
	})
}

func userSends(inputs ...string) io.Reader {
	s := ""
	for _, input := range inputs {
		s += input + "\n"
	}
	return strings.NewReader(s)
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	got := stdout.String()
	want := strings.Join(messages, "")
	if got != want {
		t.Errorf("got %q sent to user, but want %+v", got, messages)
	}
}
