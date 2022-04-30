package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/nahcnuj/til/learn-go-with-tests/my-app"
)

const dbFileName = "db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := app.NewFileSystemPlayerStore(db)
	server := app.NewServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
