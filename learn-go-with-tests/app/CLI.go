package app

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
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
