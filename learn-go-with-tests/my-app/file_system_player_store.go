package my_app

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
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
		}
	}
	return wins
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	league := s.GetLeague()

	for i, player := range league { // player is a copy of the element
		if player.Name == name {
			league[i].Wins++
		}
	}

	s.database.Seek(0, io.SeekStart)
	json.NewEncoder(s.database).Encode(league)
}
