package app

import "time"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewTexasHoldem(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{store, alerter}
}

func (g *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrementTime := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 400, 600, 1000, 2000, 4000, 8000, 16000, 32000, 64000}
	blindTime := 0 * time.Minute
	for _, amount := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, amount)
		blindTime += blindIncrementTime
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
