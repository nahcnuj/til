package my_app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (s *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.recordWin(w, player)
	}
}

func NewServer(store PlayerStore) *PlayerServer {
	s := new(PlayerServer)
	s.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandler))

	s.Handler = router

	return s
}

func (s *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := s.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) recordWin(w http.ResponseWriter, player string) {
	s.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (s *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}
}

type Player struct {
	Name string
	Wins int
}
