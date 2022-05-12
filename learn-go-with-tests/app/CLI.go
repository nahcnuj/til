package app

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputError = "Bad value for number of players"
const BadWinnerInputError = "could not recognize a winner"

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
		fmt.Fprint(cli.out, BadPlayerInputError)
		return
	}

	cli.game.Start(numberOfPlayers)

	userInput := cli.readLine()
	if !strings.HasSuffix(userInput, " wins") {
		fmt.Fprint(cli.out, BadWinnerInputError)
		return
	}
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
