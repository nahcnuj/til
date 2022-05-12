package app

import "time"

type Game struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewGame(store PlayerStore, alerter BlindAlerter) *Game {
	return &Game{store, alerter}
}

func (g *Game) Start(numberOfPlayers int) {
	blindIncrementTime := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 400, 600, 1000, 2000, 4000, 8000, 16000, 32000, 64000}
	blindTime := 0 * time.Minute
	for _, amount := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, amount)
		blindTime += blindIncrementTime
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}
