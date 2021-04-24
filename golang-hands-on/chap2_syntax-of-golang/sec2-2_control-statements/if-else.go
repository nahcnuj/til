package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type a number")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	} else {
		fmt.Println(err)
	}
}

func input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}
