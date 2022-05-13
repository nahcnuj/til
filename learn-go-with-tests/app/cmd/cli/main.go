package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nahcnuj/til/learn-go-with-tests/app"
)

const dbFileName = "db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type '{Name} wins' to record a win")

	store, close, err := app.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("could not create player store from the file, %v", err)
	}
	defer close()

	game := app.NewTexasHoldem(store, app.BlindAlerterFunc(app.Alerter))
	cli := app.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
