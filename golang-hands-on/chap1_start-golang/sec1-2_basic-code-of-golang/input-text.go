package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := input("type your name")
	fmt.Println("Hello, " + name + "!!")
}

func input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}
