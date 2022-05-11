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
	cli.in.Scan()
	cli.store.RecordWin(extractWinner(cli.in.Text()))
}

func extractWinner(input string) string {
	return strings.TrimSuffix(input, " wins")
}
