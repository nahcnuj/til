package my_app

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(s.database).Decode(&league)
	return league
}
