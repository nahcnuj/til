package app

import "sync"

type InMemoryPlayerStore struct {
	score map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{score: map[string]int{}}
}

func (s *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return s.score[player]
}

func (s *InMemoryPlayerStore) RecordWin(player string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.score[player]++
}

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.score {
		league = append(league, Player{name, wins})
	}
	return league
}
