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
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{bufio.NewScanner(in), out, game}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprintf(cli.out, "you're so silly")
		return
	}

	cli.game.Start(numberOfPlayers)

	userInput := cli.readLine()
	winner := extractWinner(userInput)
	cli.game.Finish(winner)
}

func extractWinner(input string) string {
	return strings.TrimSuffix(input, " wins")
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
