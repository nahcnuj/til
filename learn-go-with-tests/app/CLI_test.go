package app_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

func TestCLI(t *testing.T) {
	dummyPlayerStore := &app.StubPlayerStore{}
	dummyBlindAlerter := &SpyBlindAlerter{}
	dummyStdIn := &bytes.Buffer{}
	dummyStdOut := &bytes.Buffer{}

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCLI(store, in, dummyStdOut, dummyBlindAlerter)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCLI(store, in, dummyStdOut, dummyBlindAlerter)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Cleo")
	})

	t.Run("schedule printing of blind values for 5 players", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		store := &app.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := app.NewCLI(store, in, dummyStdOut, blindAlerter)
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
		cli := app.NewCLI(dummyPlayerStore, dummyStdIn, stdout, dummyBlindAlerter)
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
		cli := app.NewCLI(dummyPlayerStore, in, stdout, blindAlerter)
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

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()

	if got.amount != want.amount {
		t.Errorf("wrong blind amount, got %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("wrong time scheduled at, got %v, want %v", got.at, want.at)
	}
}
