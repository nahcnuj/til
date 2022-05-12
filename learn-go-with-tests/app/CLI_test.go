package app_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

func TestCLI(t *testing.T) {
	dummyPlayerStore := &app.StubPlayerStore{}
	dummyBlindAlerter := &SpyBlindAlerter{}
	dummyStdIn := &bytes.Buffer{}
	dummyStdOut := &bytes.Buffer{}

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		store := &app.StubPlayerStore{}
		game := app.NewGame(store, dummyBlindAlerter)

		cli := app.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		store := &app.StubPlayerStore{}
		game := app.NewGame(store, dummyBlindAlerter)

		cli := app.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Cleo")
	})

	t.Run("schedule printing of blind values for 5 players", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		store := &app.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := app.NewGame(store, blindAlerter)

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
		stdout := &bytes.Buffer{}
		game := app.NewGame(dummyPlayerStore, dummyBlindAlerter)
		cli := app.NewCLI(dummyStdIn, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := app.PlayerPrompt

		if got != want {
			t.Errorf("wrong prompt, got %q, want %q", got, want)
		}
	})

	t.Run("schedule printing of blind values for 7 players", func(t *testing.T) {
		in := strings.NewReader("7\n")
		stdout := &bytes.Buffer{}
		blindAlerter := &SpyBlindAlerter{}
		game := app.NewGame(dummyPlayerStore, blindAlerter)

		cli := app.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := app.PlayerPrompt

		if got != want {
			t.Errorf("wrong prompt, got %q, want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Minute, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 400},
			{36 * time.Minute, 600},
			{48 * time.Minute, 1000},
			{60 * time.Minute, 2000},
		}
		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert #%d was not scheduled, %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}
