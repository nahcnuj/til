package app_test

import (
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

var dummyPlayerStore = &app.StubPlayerStore{}
var dummyBlindAlerter = &SpyBlindAlerter{}

func TestTexasHoldem_Start(t *testing.T) {
	t.Run("schedule printing of blind values for 5 players", func(t *testing.T) {
		store := &app.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := app.NewTexasHoldem(store, blindAlerter)
		game.Start(5)

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
		assertScheduledAlerts(t, cases, blindAlerter.alerts)
	})

	t.Run("schedule printing of blind values for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := app.NewTexasHoldem(dummyPlayerStore, blindAlerter)
		game.Start(7)

		cases := []scheduledAlert{
			{0 * time.Minute, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 400},
			{36 * time.Minute, 600},
			{48 * time.Minute, 1000},
			{60 * time.Minute, 2000},
		}
		assertScheduledAlerts(t, cases, blindAlerter.alerts)
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

func assertScheduledAlerts(t testing.TB, cases, alerts []scheduledAlert) {
	t.Helper()

	for i, want := range cases {
		if len(alerts) <= i {
			t.Fatalf("alert #%d was not scheduled, %v", i, alerts)
		}

		got := alerts[i]
		assertScheduledAlert(t, got, want)
	}
}

func TestTexasHoldem_Finish(t *testing.T) {
	store := &app.StubPlayerStore{}
	winner := "Chris"

	game := app.NewTexasHoldem(store, dummyBlindAlerter)
	game.Finish(winner)

	app.AssertPlayerWin(t, store, winner)
}
