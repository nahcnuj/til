package main

import (
	"countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
