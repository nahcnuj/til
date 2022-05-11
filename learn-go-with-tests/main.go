package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

const dbFileName = "db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := app.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem loading to player store from file, %v", err)
	}

	server := app.NewServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
