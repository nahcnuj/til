package my_app

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	s.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(s.database)
	return league
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range s.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
