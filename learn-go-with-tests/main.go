package main

import (
	"log"
	"net/http"

	app "github.com/nahcnuj/til/learn-go-with-tests/my-app"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return 123
}

func main() {
	server := app.NewServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
