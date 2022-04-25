package my_app

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.showScore(w, r)
	case http.MethodPost:
		s.recordWin(w)
	}
}

func NewServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := s.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) recordWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}
