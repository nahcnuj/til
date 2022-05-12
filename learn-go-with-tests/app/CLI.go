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
	cli.alerter.ScheduleAlertAt(5*time.Second, 100)

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
