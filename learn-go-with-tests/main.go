package main

import (
	"log"
	"net/http"

	app "github.com/nahcnuj/til/learn-go-with-tests/my-app"
)

func main() {
	handler := http.HandlerFunc(app.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
