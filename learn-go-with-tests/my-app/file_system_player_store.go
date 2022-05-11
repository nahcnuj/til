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

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, io.SeekStart)
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem occurred loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
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
