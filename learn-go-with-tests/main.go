package main

import (
	"countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{
		Duration:  1 * time.Second,
		SleepImpl: time.Sleep,
	}
	countdown.Countdown(os.Stdout, sleeper)
}
