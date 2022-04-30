package my_app

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(s.database)
	return league
}
