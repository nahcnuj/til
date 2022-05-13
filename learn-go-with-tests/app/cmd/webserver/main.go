package main

import (
	"log"
	"net/http"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

const dbFileName = "db.json"

func main() {
	store, close, err := app.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem loading to player store from file, %v", err)
	}
	defer close()

	server, err := app.NewServer(store)
	if err != nil {
		log.Fatal("could not start a server")
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
