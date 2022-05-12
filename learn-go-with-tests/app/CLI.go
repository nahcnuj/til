package app

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
	out   io.Writer
	game  *Game
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{store, bufio.NewScanner(in), out, &Game{alerter}}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)

	numberOfPlayers, _ := strconv.Atoi(cli.readLine())

	cli.game.Start(numberOfPlayers)

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
