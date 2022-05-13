package app

import (
	"io"
	"testing"
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

func AssertGameStartedWith(t testing.TB, game *SpyGame, want int) {
	t.Helper()
	if game.StartedWith != want {
		t.Errorf("expected %d players, but got %d", want, game.StartedWith)
	}
}

func AssertGameFinishedWith(t testing.TB, game *SpyGame, want string) {
	t.Helper()
	if game.FinishedWith != want {
		t.Errorf("expected winner %s, but got %q", want, game.FinishedWith)
	}
}

func AssertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.CalledStart {
		t.Error("game should not have started")
	}

}
