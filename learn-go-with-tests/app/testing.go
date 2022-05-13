package app

import (
	"io"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin", len(store.winCalls))
	}
	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner, got %q, want %q", store.winCalls[0], winner)
	}
}

type SpyGame struct {
	CalledStart bool
	StartedWith int

	BlindAlert []byte

	CalledFinish bool
	FinishedWith string
}

func (g *SpyGame) Start(numberOfPlayers int, alertDestination io.Writer) {
	g.CalledStart = true
	g.StartedWith = numberOfPlayers

	alertDestination.Write(g.BlindAlert)
}

func (g *SpyGame) Finish(winner string) {
	g.CalledFinish = true
	g.FinishedWith = winner
}

const timeout = 500 * time.Millisecond

func AssertGameStartedWith(t testing.TB, game *SpyGame, numberOfPlayers int) {
	t.Helper()
	passed := retryUntil(timeout, func() bool {
		return game.StartedWith == numberOfPlayers
	})
	if !passed {
		t.Errorf("expected start called with %d players, but got %d", numberOfPlayers, game.StartedWith)
	}
}

func AssertGameFinishedWith(t testing.TB, game *SpyGame, winner string) {
	t.Helper()
	passed := retryUntil(timeout, func() bool {
		return game.FinishedWith == winner
	})
	if !passed {
		t.Errorf("expected finish called with %s, but got %q", winner, game.FinishedWith)
	}
}

func AssertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.CalledStart {
		t.Error("game should not have started")
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
