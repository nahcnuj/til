package my_app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	database.Seek(0, io.SeekStart)
	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("problem occurred loading player store from file %s, %v", database.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{database}),
		league:   league,
	}, nil
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

	s.database.Encode(s.league)
}
