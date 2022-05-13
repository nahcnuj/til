package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
	GetLeague() League
}

const htmlTemplatePath = "game.html"

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewServer(store PlayerStore) *PlayerServer {
	s := new(PlayerServer)
	s.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandler))
	router.Handle("/game", http.HandlerFunc(s.gameHandler))
	router.Handle("/ws", http.HandlerFunc(s.wsHandler))

	s.Handler = router

	return s
}

const jsonContentType = "application/json"

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.store.GetLeague())
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

func (s *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *PlayerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	s.store.RecordWin(string(winnerMsg))
}

type Player struct {
	Name string
	Wins int
}
