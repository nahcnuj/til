package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	appendText := func(f *os.File, s string) {
		_, err := f.WriteString(s + "\n")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			f.Close()
			return
		}
	}

	f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	appendText(f, "*** start ***")
	for {
		s := input("type message")
		if s == "" {
			break
		}
		appendText(f, s)
	}
	appendText(f, "*** end ***")

	if err = f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}
