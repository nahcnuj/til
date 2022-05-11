package app

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store, in}
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.store.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(input string) string {
	return strings.TrimSuffix(input, " wins")
}
