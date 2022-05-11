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

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("could not open %s, %v", dbFileName, err)
	}

	store, err := app.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("could not create player store from the file, %v", err)
	}

	game := app.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
