package app_test

import (
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
	dummySpyAlerter := &SpyBlindAlerter{}

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		app.AssertPlayerWin(t, store, "Cleo")
	})

	t.Run("schedule printing of blind values for 5 players", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &app.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := app.NewCLI(store, in, blindAlerter)
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

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.at), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert #%d was not scheduled, %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]

				gotAmount := alert.amount
				if gotAmount != c.amount {
					t.Errorf("wrong blind amount, got %d, want %d", gotAmount, c.amount)
				}

				gotScheduledTime := alert.at
				if gotScheduledTime != c.at {
					t.Errorf("wrong time scheduled at, got %v, want %v", gotScheduledTime, c.at)
				}
			})
		}
	})
}
