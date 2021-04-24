package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := input("type a price without tax")
	price, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error!")
		return
	}
	fmt.Println(int(float64(price) * 1.1))
}

func input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}
