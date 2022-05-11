package app

import "io"

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.store.RecordWin("Cleo")
}
