package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{store, bufio.NewScanner(in), out, alerter}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)

	cli.scheduleBlindAlerts()

	userInput := cli.readLine()
	cli.store.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 400, 600, 1000, 2000, 4000, 8000, 16000, 32000, 64000}
	blindTime := 0 * time.Minute
	for _, amount := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, amount)
		blindTime += 10 * time.Minute
	}
}

func extractWinner(input string) string {
	return strings.TrimSuffix(input, " wins")
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
