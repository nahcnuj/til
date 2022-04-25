package main

import (
	"log"
	"net/http"

	app "github.com/nahcnuj/til/learn-go-with-tests/my-app"
)

func main() {
	server := &app.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
