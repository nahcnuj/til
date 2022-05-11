package my_app

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: &tape{database},
		league:   league,
	}
}

func (s *FileSystemPlayerStore) GetLeague() League {
	return s.league
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := s.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	player := s.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{name, 1})
	}

	json.NewEncoder(s.database).Encode(s.league)
}
