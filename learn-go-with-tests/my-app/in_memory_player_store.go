package my_app

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return 123
}

func (s *InMemoryPlayerStore) RecordWin(player string) {

}
