package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	template *template.Template
	game     Game
}

func NewServer(store PlayerStore, game Game) (*PlayerServer, error) {
	s := new(PlayerServer)
	s.store = store
	s.game = game

	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem loading template: %s", err.Error())
	}
	s.template = tmpl

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandler))
	router.Handle("/game", http.HandlerFunc(s.gameHandler))
	router.Handle("/ws", http.HandlerFunc(s.wsHandler))

	s.Handler = router

	return s, nil
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
	s.template.Execute(w, nil)
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type playerServerWS struct {
	*websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to WebSocket: %v", err)
	}
	return &playerServerWS{conn}
}

func (s *playerServerWS) WaitForMsg() string {
	_, msg, err := s.ReadMessage()
	if err != nil {
		log.Printf("problem reading message from WebSocket: %v", err)
	}
	return string(msg)
}

func (s *playerServerWS) Write(p []byte) (n int, err error) {
	err = s.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (s *PlayerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	s.game.Start(numberOfPlayers, ws)

	winner := ws.WaitForMsg()
	s.game.Finish(string(winner))
}

type Player struct {
	Name string
	Wins int
}
