package app

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	alerter BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{store, bufio.NewScanner(in), alerter}
}

func (cli *CLI) PlayPoker() {
	blinds := []int{100, 200, 400, 600, 1000, 2000, 4000, 8000, 16000, 32000, 64000}
	blindTime := 0 * time.Minute
	for _, amount := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, amount)
		blindTime += 10 * time.Minute
	}

	userInput := cli.readLine()
	cli.store.RecordWin(extractWinner(userInput))
}

func extractWinner(input string) string {
	return strings.TrimSuffix(input, " wins")
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
