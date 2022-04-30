package my_app

import "io"

type FileSystemPlayerStore struct {
	database io.Reader
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	return nil
}
