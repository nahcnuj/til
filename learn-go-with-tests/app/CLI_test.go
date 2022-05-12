package app_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

var dummyStdOut = &bytes.Buffer{}

type SpyGame struct {
	StartedWith  int
	FinishedWith string
}

func (g *SpyGame) Start(numberOfPlayers int) {
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

	t.Run("schedule printing of blind values for 5 players", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		store := &app.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := app.NewTexasHoldem(store, blindAlerter)

		cli := app.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Minute, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 400},
			{30 * time.Minute, 600},
			{40 * time.Minute, 1000},
			{50 * time.Minute, 2000},
			{60 * time.Minute, 4000},
			{70 * time.Minute, 8000},
			{80 * time.Minute, 16000},
			{90 * time.Minute, 32000},
			{100 * time.Minute, 64000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", want.amount, want.at), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert #%d was not scheduled, %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
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

		if game.StartedWith != 7 {
			t.Errorf("expected 7 players but got %d", game.StartedWith)
		}
	})
}

func assertWinner(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected winner %s, but got %q", want, got)
	}
}
